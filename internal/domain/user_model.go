package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint    `gorm:"primarykey"`
	Nickname  string  `gorm:"column:nickname;type:varchar(255);unique"` // min(2), max(12)
	Biography *string `gorm:"column:biography;type:varchar(255);"`      // min(0), max(14)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Groups Groups `gorm:"many2many:user_group_mappers;"`
	Goals  Goals
	Topics Topics
}

type CreateUserParams = User

type ListUsersParams struct {
	// conditions
	IDs         []uint
	Nicknames   []string
	Biographies []string

	// order
	Order   SortOrder
	OrderBy string

	// relation
	WithGroups      bool
	WithGroupsLimit int
	WithGroupsIDs   []uint

	WithGoals  bool
	WithTopics bool

	// options
	Limit  int
	Offset int
}

type PatchUserParams struct {
	// conditions
	ID uint

	// updates
	Nickname  *string
	Biography *string
}

type DeleteUserParams struct {
	// conditions
	ID       uint
	Nickname string

	// option
	IsHardDelete bool
}

type AppendGroupsParams struct {
	GroupIDs []uint
	UserID   uint
}

type RemoveGroupsParams struct {
	UserID   uint
	GroupIDs []uint
}

type FetchGroupsParams struct {
	Limit  int
	UserID uint
}

type HasGroupsParams struct {
	UserID   uint
	GroupIDs []uint
}
