package service

import (
	"context"
	"errors"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
	"github.com/samber/lo"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.CreateUserResponse{}, domain.ErrUserAlreadyExists
		}

		return domain.CreateUserResponse{}, err
	}

	return domain.CreateUserResponse{
		UserID:    createdUser.ID,
		Nickname:  createdUser.Nickname,
		Biography: lo.FromPtr(createdUser.Biography),
	}, nil
}

func (s UserService) PatchUser(ctx context.Context, request domain.PatchUserRequest) error {
	// (1) 수정하고자 하는 사용자가 존재하는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs: []uint{request.RequestUserID},
	})
	if err != nil {
		return err
	}
	if foundUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}

	// (2) 사용자 수정
	err = s.userStore.PatchUser(ctx, domain.PatchUserParams{
		ID:        request.RequestUserID,
		Nickname:  request.Nickname,
		Biography: request.Biography,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) DeleteUser(ctx context.Context, request domain.DeleteUserRequest) error {
	// (1) 삭제하려고 하는 사용자가 존재하는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}

	// (2) 사용자 삭제
	err = s.userStore.DeleteUser(ctx, domain.DeleteUserParams{
		ID: request.RequestUserID,
	})
	if err != nil {
		return err
	}

	return nil
}
