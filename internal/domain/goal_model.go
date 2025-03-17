package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Page     int       `gorm:"column:page;type:integer"`
	Deadline time.Time `gorm:"column:deadline;type:timestamp"`

	UserID  uint
	GroupID uint
}
