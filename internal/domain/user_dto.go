package domain

type CreateUserRequest struct {
	Nickname  string  `json:"nickname"  validate:"required,min=2,max=12"`
	Biography *string `json:"biography" validate:"max=14"`
}

type CreateUserResponse struct {
	UserID    uint   `json:"userID"`
	Nickname  string `json:"nickname"`
	Biography string `json:"biography"`
}

type CreateRandomUserResponse struct {
	UserID    uint   `json:"userID"`
	Nickname  string `json:"nickname"`
	Biography string `json:"biography"`
}

type PatchUserRequest struct {
	// conditions
	RequestUserID uint `json:"-" validate:"required,gte=1"`

	// updates
	Nickname  *string `json:"nickname"  validate:"omitempty,min=2,max=12"`
	Biography *string `json:"biography" validate:"max=14"`
}

type DeleteUserRequest struct {
	// conditions
	RequestUserID uint `json:"-" validate:"required,gte=1"`
}
