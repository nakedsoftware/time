package database

import (
	"database/sql"
	"time"
)

type Pomodoro struct {
	Model

	StartTime time.Time
	EndTime   sql.NullTime
	Completed bool
}
