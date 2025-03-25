package domain

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"column:title;type:varchar(255)"`   // min(4) max(64)
	Content   string `gorm:"column:content;type:varchar(255)"` // min(4) max(128)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GoalID uint
	UserID uint
}

type CreateTopicParams = Topic

type ListTopicsParams struct {
	// conditions
	IDs      []uint
	Titles   []string
	Contents []string

	// order
	Order   SortOrder
	OrderBy string

	// options
	Limit  int
	Offset int
}

type PatchTopic struct {
	// condition
	ID uint

	// updates
	Title   *string
	Content *string
}

type DeleteTopic struct {
	// condition
	ID uint

	// options
	IsHardDelete bool
}
