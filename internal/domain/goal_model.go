package domain

import "gorm.io/gorm"

type Goal struct {
	gorm.Model

	UserID uint
}
