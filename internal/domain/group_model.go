package domain

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Group struct {
	ID              uint    `gorm:"primarykey"`
	Introduction    *string `gorm:"column:introduction;type:varchar(255)"`           // min(0), max(255)
	BookTitle       string  `gorm:"column:book_title;type:varchar(255)"`             // min(1), max(255)
	BookAuthor      string  `gorm:"column:book_author;type:varchar(255);null"`       // min(1), max(255)
	BookPublisher   *string `gorm:"column:book_publisher;type:varchar(255);null"`    // min(0), max(255)
	BookMaxPage     int     `gorm:"column:book_max_page;type:integer"`               // gte(1)
	BookCurrentPage int     `gorm:"column:book_current_page;type:integer;default:0"` // gte(1)
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Users Users `gorm:"many2many:user_group_mappers;"`
	Goals Goals
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
	WithUsers    bool
	WithUsersIDs []uint

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
