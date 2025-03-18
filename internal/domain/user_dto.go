package domain

import "github.com/samber/lo"

type CreateUserRequest struct {
	Nickname  string  `json:"nickname" validate:"required"`
	Biography *string `json:"biography" validate:"omitempty,max=14"`
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

type ListUsersRequest struct {
	UserIDs     []uint   `query:"userIDs"`
	Nicknames   []string `query:"nicknames"`
	Biographies []string `query:"biographies" `
}

type ListUsersResponse struct {
	Users []UsersResponse `json:"users"`
}

type UsersResponse struct {
	UserID    uint   `json:"userID"`
	Nickname  string `json:"nickname"`
	Biography string `json:"biography"`
}

func (m Users) ToListUsersResponse() ListUsersResponse {
	usersResponse := make([]UsersResponse, 0, len(m))

	for _, user := range m {
		usersResponse = append(usersResponse, UsersResponse{
			UserID:    user.ID,
			Nickname:  user.Nickname,
			Biography: lo.FromPtr(user.Biography),
		})
	}

	listUsersResponse := ListUsersResponse{
		Users: usersResponse,
	}

	return listUsersResponse
}

type PatchUserRequest struct {
	// conditions
	UserID uint `json:"-"`

	// updates
	Nickname  *string `json:"nickname" validate:"omitempty"`
	Biography *string `json:"biography" validate:"omitempty"`
}

type DeleteUserRequest struct {
	// conditions
	UserID uint `json:"-"`
}
