package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/nakedsoftware/time/internal/database"
	"github.com/nakedsoftware/time/internal/pomodoro"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCommand = &cobra.Command{
	Use:     "time",
	Version: "0.0.1",
	Short:   "Naked Time is a time management and time tracking application.",
	Long: `
Naked Time is a time management and time tracking application that helps you
to focus on completing important tasks and analyzing how you are using your
time. Naked Time implements multiple tools to help you to improve your time
management skills and productivity. Using Naked Time, you can achieve better
focus in managing and tracking how much time you are working on activities,
using your time to focus on achieving value and completing important
activities, and analyzing if you are using your time effectively or can make
improvements in your estimation and time management skills.
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		homePath, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		dbDir := path.Join(homePath, ".nakedtime")
		dbPath := path.Join(dbDir, "time.db")
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return err
		}

		db, err := database.NewDB(dbPath)
		if err != nil {
			return err
		}

		ctx := context.WithValue(cmd.Context(), databaseContextKey, db)
		cmd.SetContext(ctx)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		db := cmd.Context().Value(databaseContextKey).(*gorm.DB)
		id, err := startPomodoro(cmd.Context(), db)
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

func startPomodoro(
	ctx context.Context,
	db *gorm.DB,
) (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return id, err
	}

	p := &database.Pomodoro{
		Model: database.Model{
			ID: id,
		},
		StartTime: time.Now(),
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
