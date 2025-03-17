package domain

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(255)"`   // min(4) max(24)
	Content string `gorm:"column:content;type:varchar(255)"` // min(4) max(128)

	GoalID uint
	UserID uint
}
