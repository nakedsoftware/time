// Copyright 2025 Naked Software, LLC
//
// License Agreement
//
// Version 1.0.0 (September 28, 2025)
//
// This Naked Time License Agreement ("Agreement") is a legal agreement between
// you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the
// Naked Time software product ("Software"). By using the Software, you agree
// to be bound by the terms of this Agreement.
//
// 1. Grant of License
//
// Licensor grants Licensee a non-exclusive, non-transferable,
// non-sublicensable license to use the Software for non-commercial,
// educational, or other non-production purposes. Licensee may not use the
// Software for any commercial purposes without purchasing a commercial license
// from Licensor.
//
// 2. Commercial Use
//
// To use the Software for commercial purposes, Licensee must purchase a
// commercial license from Licensor. A commercial license allows Licensee to
// use the Software in production environments, build their own version, and
// add custom features or bug fixes. Licensee may not sell the Software or any
// derivative works to others.
//
// 3. Derivative Works
//
// Licensee may create derivative works of the Software for their own use,
// provided that they maintain a valid commercial license. Licensee may not
// sell or distribute derivative works to others. Any derivative works must
// include a copy of this Agreement and retain all copyright notices.
//
// 4. Sharing and Contributions
//
// Licensee may share their changes or bug fixes to the Software with others,
// provided that such changes are made freely available and not sold. Licensee
// is encouraged to contribute their bug fixes back to Licensor for inclusion
// in the Software.
//
// 5. Restrictions
//
// Licensee may not:
//
// - Use the Software for any commercial purposes without a valid commercial
//   license.
// - Sell, sublicense, or distribute the Software or any derivative works.
// - Remove or alter any copyright notices or proprietary legends on the
//   Software.
//
// 6. Termination
//
// This Agreement is effective until terminated. Licensor may terminate this
// Agreement at any time if Licensee breaches any of its terms. Upon
// termination, Licensee must cease all use of the Software and destroy all
// copies in their possession.
//
// 7. Disclaimer of Warranty
//
// The Software is provided "as is" without warranty of any kind, express or
// implied, including but not limited to the warranties of merchantability,
// fitness for a particular purpose, and noninfringement. In no event shall
// Licensor be liable for any claim, damages, or other liability, whether in
// an action of contract, tort, or otherwise, arising from, out of, or in
// connection with the Software or the use or other dealings in the Software.
//
// 8. Limitation of Liability
//
// In no event shall Licensor be liable for any indirect, incidental, special,
// exemplary, or consequential damages (including, but not limited to,
// procurement of substitute goods or services; loss of use, data, or profits;
// or business interruption) however caused and on any theory of liability,
// whether in contract, strict liability, or tort (including negligence or
// otherwise) arising in any way out of the use of the Software, even if
// advised of the possibility of such damage.
//
// 9. Governing Law
//
// This Agreement shall be governed by and construed in accordance with the
// laws of the jurisdiction in which Licensor is located, without regard to its
// conflict of law principles.
//
// 10. Entire Agreement
//
// This Agreement constitutes the entire agreement between the parties with
// respect to the Software and supersedes all prior or contemporaneous
// understandings regarding such subject matter.
//
// By using the Software, you acknowledge that you have read this Agreement,
// understand it, and agree to be bound by its terms and conditions.

package pomodoro

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	appcontext "github.com/nakedsoftware/time/internal/context"
	"github.com/nakedsoftware/time/internal/database"
	"github.com/nakedsoftware/time/internal/pomodoro"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var StartCommand = &cobra.Command{
	Use:   "start activity-id",
	Short: "Starts a pomodoro",
	Long: `
The start command will start a pomodoro to allow you to focus on completing
an important activity. The command will present a timer for 25 minutes during
which you can focus on completing the work in front of you. After the pomodoro
completes, an alarm will sound and the pomodoro will be recorded as being
completed.
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			activityIDStr string
			err           error
		)
		if len(args) > 0 {
			activityIDStr = args[0]
		} else {
			activityIDStr, err = readActivityID(cmd.Context())
			if err != nil {
				return err
			}
		}

		activityID, err := uuid.Parse(activityIDStr)
		if err != nil {
			return err
		}

		db := appcontext.GetDB(cmd)

		id, err := startPomodoro(cmd.Context(), db, activityID)
		if err != nil {
			return err
		}

		completed, err := pomodoro.Run(cmd.Context())
		if err != nil {
			return err
		}

		return endPomodoro(cmd.Context(), db, id, completed)
	},
}

func readActivityID(ctx context.Context) (string, error) {
	// readActivityID will attempt to read the activity ID from stdin. A
	// timeout context is used to avoid blocking indefinitely if no data is
	// available on stdin.

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	done := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			done <- strings.TrimSpace(scanner.Text())
		} else {
			if err := scanner.Err(); err != nil {
				errChan <- err
			} else {
				errChan <- io.EOF
			}
		}
	}()

	select {
	case activityIDStr := <-done:
		return activityIDStr, nil

	case err := <-errChan:
		if err == io.EOF {
			return "", fmt.Errorf("no input provided for activity ID (EOF)")
		}

		return "", err

	case <-ctx.Done():
		return "", fmt.Errorf("timed out waiting for activity ID input")
	}
}

func startPomodoro(
	ctx context.Context,
	db *gorm.DB,
	activityID uuid.UUID,
) (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return id, err
	}

	p := &database.Pomodoro{
		Model: database.Model{
			ID: id,
		},
		ActivityID: activityID,
		StartTime:  time.Now(),
	}
	err = gorm.G[database.Pomodoro](db).Create(ctx, p)
	return id, err
}

func endPomodoro(
	ctx context.Context,
	db *gorm.DB,
	id uuid.UUID,
	completed bool,
) error {
	rowsAffected, err := gorm.G[database.Pomodoro](db).
		Where("id = ?", id).
		Updates(
			ctx,
			database.Pomodoro{
				EndTime: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				Completed: completed,
			},
		)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf(
			"pomodoro not found in the database to be closed",
		)
	}

	return nil
}
