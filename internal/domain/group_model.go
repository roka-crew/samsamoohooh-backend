package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Introduction    string
	BookTitle       string
	BookAuthor      string
	BookPublisher   string
	BookMaxPage     int
	BookCurrentPage int

	Users []User
}
