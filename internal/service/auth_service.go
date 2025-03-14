package service

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/token"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
)

type AuthService struct {
	userStore *store.UserStore
	jwtMaker  *token.JWTMaker
}

func NewAuthService(
	userStore *store.UserStore,
	jwtMaker *token.JWTMaker,
) *AuthService {
	return &AuthService{
		userStore: userStore,
		jwtMaker:  jwtMaker,
	}
}

func (s AuthService) Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, error) {
	// (1) 존재하는 사용자인이 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		Nicknames: []string{request.Nickname},
		Limit:     1,
	})
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// (2) 해당 사용자의 토큰 생성
	createdTokenString, err := s.jwtMaker.CreateTokenString(foundUsers.First().ID)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		Token: createdTokenString,
	}, nil
}
