package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Introduction    *string `gorm:"column:introduction;type:varchar(255)"`
	BookTitle       string  `gorm:"column:book_title;type:varchar(255)"`
	BookAuthor      string  `gorm:"column:book_author;type:varchar(255)"`
	BookPublisher   string  `gorm:"column:book_publisher;type:varchar(255)"`
	BookMaxPage     int     `gorm:"column:book_max_page;type:integer"`
	BookCurrentPage int     `gorm:"column:book_current_page;type:integer;default:0"`

	Users []User `gorm:"many2many:user_group_mappers;"`
	Goals []Goal
}
