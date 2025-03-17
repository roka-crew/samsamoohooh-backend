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

type CreateGroupParams = Group

type ListGroupsParams struct {
	// conditions
	IDs              []uint
	BookTitles       []string
	BookAuthors      []string
	BookPublishers   []string
	BookMaxPages     []int
	BookCurrentPages []int

	// order
	Order   SortOrder
	OrderBy string

	// relation
	WithUsers bool
	WithGoals bool

	// options
	Limit  int
	Offset int
}

type PatchGroupParams struct {
	// conditions
	ID uint

	// udpates
	Introduction    *string
	BookTitle       *string
	BookAuthor      *string
	BookPublisher   *string
	BookMaxPage     *int
	BookCurrentPage *int
}

type DeleteGroupParams struct {
	// conditions
	ID uint

	// option
	IsHardDelete bool
}

type AppendUserParams struct {
	GroupID uint
	UserIDs []uint
}

type RemoveUsersParams struct {
	GroupID uint
	UserIDs []uint
}

type FetchUsersParams struct {
	GroupIDs []uint
}
