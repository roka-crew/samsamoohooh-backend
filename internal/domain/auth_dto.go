package domain

type LoginRequest struct {
	Nickname string `json:"nickname" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ValidateRequest struct {
	BearerToken string `reqHeader:"Authorization"`
}

type ValidateResponse struct {
	UserID    uint    `json:"userID"`
	Nickname  string  `json:"nickname"`
	Biography *string `json:"biography"`
}
