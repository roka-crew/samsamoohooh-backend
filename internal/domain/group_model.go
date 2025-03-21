package domain

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Introduction    *string `gorm:"column:introduction;type:varchar(255)"`
	BookTitle       string  `gorm:"column:book_title;type:varchar(255)"`
	BookAuthor      string  `gorm:"column:book_author;type:varchar(255);null"`
	BookPublisher   *string `gorm:"column:book_publisher;type:varchar(255);null"`
	BookMaxPage     int     `gorm:"column:book_max_page;type:integer"`
	BookCurrentPage int     `gorm:"column:book_current_page;type:integer;default:0"`

	Users []User `gorm:"many2many:user_group_mappers;"`
	Goals []Goal
}

func (g Group) ToCreateGroupResponse() CreateGroupResponse {
	return CreateGroupResponse{
		GroupID:         g.ID,
		Introduction:    lo.FromPtr(g.Introduction),
		BookTitle:       g.BookTitle,
		BookAuthor:      g.BookAuthor,
		BookPublisher:   lo.FromPtr(g.BookPublisher),
		BookMaxPage:     g.BookMaxPage,
		BookCurrentPage: g.BookCurrentPage,
	}
}

func (g Groups) ToListGroupsResponse() ListGroupsResponse {
	groupsResponse := make([]GroupResponse, 0, len(g))

	for _, group := range g {
		groupsResponse = append(groupsResponse, GroupResponse{
			GroupID:         group.ID,
			BookTitle:       group.BookTitle,
			BookAuthor:      group.BookAuthor,
			BookPublisher:   lo.FromPtr(group.BookPublisher),
			BookMaxPage:     group.BookMaxPage,
			BookCurrentPage: group.BookCurrentPage,
			Introduction:    lo.FromPtr(group.Introduction),
		})
	}

	return ListGroupsResponse{Groups: groupsResponse}
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
	Limit    int
}
