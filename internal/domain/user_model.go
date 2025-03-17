package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname  string  `gorm:"column:nickname;type:varchar(255);unique"` // min(2), max(12)
	Biography *string `gorm:"column:biography;type:varchar(255);"`      // min(0), max(14)

	Groups []Group `gorm:"many2many:user_group_mappers;"`
	Goals  []Goal
	Topics []Topic
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
	WithGroups bool
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
	UserID   uint
	GroupIDs []uint
}

type RemoveGroupsParams struct {
	UserID   uint
	GroupIDs []uint
}

type FetchGroupsParams struct {
	UserIDs []uint
}
