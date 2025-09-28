package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model

	ID uuid.UUID `gorm:"type:text;primaryKey"`
}
