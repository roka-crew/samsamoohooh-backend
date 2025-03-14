package domain

type LoginRequest struct {
	Nickname string `json:"nickname" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
