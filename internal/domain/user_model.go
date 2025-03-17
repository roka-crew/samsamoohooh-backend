package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname  string  `gorm:"column:nickname;type:varchar(255);unique"` // min(2), max(12)
	Biography *string `gorm:"column:biography;type:varchar(255);"`      // min(0), max(14)

	Groups []Group
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
