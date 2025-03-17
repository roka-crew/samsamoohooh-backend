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

type CreateGoalParams = Goal

type ListGoalsParmas struct {
	// conditions
	IDs       []uint
	Pages     []int
	Deadlines []time.Time

	// order
	Order   SortOrder
	OrderBy string

	// options
	Limit  int
	Offset int
}

type PatchGoalParams struct {
	// condition
	ID uint

	// updates
	Page     *int
	Deadline *time.Time
}

type DeleteGoalParams struct {
	// condition
	ID uint

	// options
	IsHardDelete bool
}
