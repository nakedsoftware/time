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

package activity

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/nakedsoftware/time/internal/activities"
	appcontext "github.com/nakedsoftware/time/internal/context"
	"github.com/nakedsoftware/time/internal/database"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	beforeID string
	afterID  string

	OrderCommand = &cobra.Command{
		Use:   "order",
		Short: "Prioritize the Activity Inventory",
		Long: `
The activity order command is used to prioritize the activities in the
Activity Inventory. The order command is interactive and will allow you to
view the active activities in the Activity Inventory, select the activity,
and move it to its new place in the Activity Inventory.
`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				id, err := uuid.Parse(args[0])
				if err != nil {
					return err
				}

				if len(beforeID) > 0 {
					otherID, err := uuid.Parse(beforeID)
					if err != nil {
						return err
					}

					return moveActivityBefore(cmd.Context(), id, otherID)
				}

				if len(afterID) > 0 {
					otherID, err := uuid.Parse(afterID)
					if err != nil {
						return err
					}

					return moveActivityAfter(cmd.Context(), id, otherID)
				}

				return fmt.Errorf(
					"either --before or --after must be specified",
				)
			}

			return showOrderUI(cmd.Context())
		},
	}
)

func init() {
	OrderCommand.Flags().StringVar(
		&afterID,
		"after",
		"",
		"ID of the activity to move the activity after",
	)
	OrderCommand.Flags().StringVar(
		&beforeID,
		"before",
		"",
		"ID of the activity to move the activity in front of",
	)
}

func showOrderUI(ctx context.Context) error {
	p := tea.NewProgram(activities.NewModel(
		ctx,
		appcontext.GetDB(ctx),
	))
	_, err := p.Run()
	return err
}

// reorderActivities moves an activity to a new position in the activity list.
// If insertAfter is true, the activity is inserted after the target position,
// otherwise it's inserted before the target position.
func reorderActivities(
	ctx context.Context,
	id, otherID uuid.UUID,
	insertAfter bool,
) error {
	db := appcontext.GetDB(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		activityList, err := gorm.G[database.Activity](db).
			Where("completed = ?", false).
			Order("priority ASC, created_at ASC").
			Find(ctx)
		if err != nil {
			return err
		}

		// Find the indices of both activities
		var activityIndex = -1
		var otherIndex = -1

		for i, activity := range activityList {
			if activity.ID == id {
				activityIndex = i
			}
			if activity.ID == otherID {
				otherIndex = i
			}
		}

		// Validate both activities were found
		if activityIndex == -1 {
			return fmt.Errorf("activity with ID %s not found", id)
		}
		if otherIndex == -1 {
			return fmt.Errorf("activity with ID %s not found", otherID)
		}

		// Remove the activity from its current position
		activityToMove := activityList[activityIndex]
		activityList = append(
			activityList[:activityIndex],
			activityList[activityIndex+1:]...,
		)

		// Adjust otherIndex if needed (if we removed an element before it)
		if activityIndex < otherIndex {
			otherIndex--
		}

		// Calculate insert position
		insertPosition := otherIndex
		if insertAfter {
			insertPosition++
		}

		// Insert the activity at the target position
		activityList = append(
			activityList[:insertPosition],
			append(
				[]database.Activity{activityToMove},
				activityList[insertPosition:]...,
			)...,
		)

		// Update all priorities from 1 to len(activityList)
		for i := range activityList {
			activityList[i].Priority = i + 1
			if err := tx.Save(&activityList[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func moveActivityBefore(ctx context.Context, id, otherID uuid.UUID) error {
	return reorderActivities(ctx, id, otherID, false)
}

func moveActivityAfter(ctx context.Context, id, otherID uuid.UUID) error {
	return reorderActivities(ctx, id, otherID, true)
}
