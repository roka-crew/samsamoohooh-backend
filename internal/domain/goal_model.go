package domain

import (
	"time"

	"gorm.io/gorm"
)

type GoalStatus string

const (
	GoalStatusDiscussionPending GoalStatus = "DISCUSSION_PENDING"
	GoalStatusDiscussionDone    GoalStatus = "DISCUSSION_DONE"
)

type Goal struct {
	ID       uint       `gorm:"primarykey"`
	Page     int        `gorm:"column:page;type:integer"`
	Deadline time.Time  `gorm:"column:deadline;type:timestamp"`
	Status   GoalStatus `gorm:"column:status;type:varchar(255)"` // ENUM('DISCUSSION_PENDING', 'DISCUSSION_DONE')

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Topics  Topics
	UserID  uint
	GroupID uint
}

type CreateGoalParams = Goal

type ListGoalsParmas struct {
	// conditions
	IDs         []uint
	Pages       []int
	Deadlines   []time.Time
	Statuses    []GoalStatus
	GtCreatedAt []time.Time
	GroupIDs    []uint

	// order
	Order   SortOrder
	OrderBy string

	// relation
	WithTopics      bool
	WithTopicsLimit int

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
