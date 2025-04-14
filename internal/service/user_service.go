package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type UserService struct {
	userStore *store.UserStore
}

func NewUserService(
	userStore *store.UserStore,
) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (s UserService) CreateUser(ctx context.Context, request domain.CreateUserRequest) (domain.CreateUserResponse, error) {
	// (1) 사용자 Nickname 중복 여부를 검사
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

func (s UserService) CreateRandomUser(ctx context.Context) (domain.CreateRandomUserResponse, error) {
	var randomNicknames = []string{"토이스토리", "푸딩", "별똥별", "오보에", "늑대", "고래", "보니하니", "트론본", "충만", "바순"}
	randomNickname := randomNicknames[rand.IntN(len(randomNicknames))]

	// (1) 가장 최신에 등록한 사용자의 ID 가져오기
	latestUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		OrderBy: domain.ModelUserID,
		Order:   domain.SortOrderDesc,

		Limit: 1,
	})
	if err != nil {
		return domain.CreateRandomUserResponse{}, err
	}

	// (2) 식별자 정하기
	var identifier = 1
	if !latestUsers.IsEmpty() {
		identifier = int(latestUsers.First().ID) + 1
	}

	// (3) 죄총 nickname 생성
	finalNickname := fmt.Sprintf("%s%d", randomNickname, identifier)

	// (4) 새로운 사용자 생성
	createdUser, err := s.userStore.CreateUser(ctx, domain.CreateUserParams{
		Nickname: finalNickname,
	})
	if err != nil {
		return domain.CreateRandomUserResponse{}, err
	}

	return domain.CreateRandomUserResponse{
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
