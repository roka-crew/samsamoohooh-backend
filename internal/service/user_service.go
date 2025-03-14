package service

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
)

type UserService struct {
	userStore *store.UserStore
}

func NewUserSerivce(
	userStore *store.UserStore,
) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (s UserService) CreateUser(ctx context.Context, request domain.CreateUserRequest) (domain.CreateUserResponse, error) {
	// (1) 사용자 Nickname의 중복 여부를 검사
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		Nicknames: []string{request.Nickname},
	})
	if err != nil {
		return domain.CreateUserResponse{}, err
	}
	if !foundUsers.IsEmpty() {
		return domain.CreateUserResponse{}, domain.ErrUserAlreadyExists
	}

	// (2) 새로운 사용자 생성
	createdUser, err := s.userStore.CreateUser(ctx, domain.CreateUserParams{
		Nickname:  request.Nickname,
		Biography: request.Biography,
	})
	if err != nil {
		return domain.CreateUserResponse{}, err
	}

	return createdUser.ToCreateUserResponse(), nil
}

func (s UserService) ListUsers(ctx context.Context, request domain.ListUsersRequest) (domain.ListUsersResponse, error) {
	// (1) 사용자 조회
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:         request.UserIDs,
		Nicknames:   request.Nicknames,
		Biographies: request.Biographies,
	})
	if err != nil {
		return domain.ListUsersResponse{}, err
	}

	return foundUsers.ToListUsersResponse(), nil
}

func (s UserService) PatchUser(ctx context.Context, request domain.PatchUserRequest) error {
	// (1) 수정하고자 하는 사용자가 존재하는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs: []uint{request.UserID},
	})
	if err != nil {
		return domain.ErrUserNotFound
	}
	if foundUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}

	// (2) 사용자 수정
	err = s.userStore.PatchUser(ctx, domain.PatchUserParams{
		ID:        request.UserID,
		Nickname:  request.Nickname,
		Biography: request.Biography,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) DeleteUser(ctx context.Context, request domain.DeleteUserRequest) error {
	err := s.userStore.DeleteUser(ctx, domain.DeleteUserParams{
		ID: request.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}
