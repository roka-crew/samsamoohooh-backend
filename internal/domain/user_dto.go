package domain

import "github.com/samber/lo"

type CreateUserRequest struct {
	Nickname  string  `json:"nickname"  validate:"required,min=2,max=12"`
	Biography *string `json:"biography" validate:"max=14"`
}

type CreateUserResponse struct {
	UserID    uint   `json:"userID"`
	Nickname  string `json:"nickname"`
	Biography string `json:"biography"`
}

func (m User) ToCreateUserResponse() CreateUserResponse {
	return CreateUserResponse{
		UserID:    m.ID,
		Nickname:  m.Nickname,
		Biography: lo.FromPtr(m.Biography),
	}
}

type PatchUserRequest struct {
	// conditions
	RequestUserID uint `json:"-" validate:"required,gte=1"`

	// updates
	Nickname  *string `json:"nickname"  validate:"min=2,max=12"`
	Biography *string `json:"biography" validate:"max=14"`
}

type DeleteUserRequest struct {
	// conditions
	RequestUserID uint `json:"-" validate:"required,gte=1"`
}
