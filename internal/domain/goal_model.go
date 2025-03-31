package domain

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID        uint      `gorm:"primarykey"`
	Page      int       `gorm:"column:page;type:integer"`
	Deadline  time.Time `gorm:"column:deadline;type:timestamp"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID  uint
	GroupID uint
}

type CreateGoalParams = Goal

type ListGoalsParmas struct {
	// conditions
	IDs         []uint
	Pages       []int
	Deadlines   []time.Time
	GtCreatedAt []time.Time
	GroupIDs    []uint

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
