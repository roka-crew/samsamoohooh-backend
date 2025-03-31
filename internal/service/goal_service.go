package service

import (
	"context"
	"time"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
)

type GoalService struct {
	goalStore  *store.GoalStore
	userStore  *store.UserStore
	groupStore *store.GroupStore
}

func NewGoalService(
	goalStore *store.GoalStore,
	userStore *store.UserStore,
	groupStore *store.GroupStore,
) *GoalService {
	return &GoalService{
		goalStore:  goalStore,
		userStore:  userStore,
		groupStore: groupStore,
	}
}

func (s *GoalService) CreateGoal(ctx context.Context, request domain.CreateGoalRequest) (domain.CreateGoalResponse, error) {
	// Goal 생성 조건
	// Goal은 2가지의 상태를 가집니다.
	// 1. 진행중인 목표  (현재 시각이, 목표의 데드라인보다 과거임)
	// 2. 데드라인이 마감된 목표 (현재 시각이, 목표의 데드라인보다 미래임)
	// 조건 1:  진행중인 목표가 존재할 때는 새로운 목표를 생성할 수 없습니다.

	// (1) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{request.GroupID},
	})
	if err != nil {
		return domain.CreateGoalResponse{}, err
	}
	if foundUsers.IsEmpty() {
		return domain.CreateGoalResponse{}, domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.CreateGoalResponse{}, domain.ErrUserNotInGroup
	}

	// (2) 진행중인 목표가 존재하는지 확인
	foundGoals, err := s.goalStore.ListGoals(ctx, domain.ListGoalsParmas{
		IDs:         []uint{request.GroupID},
		GtCreatedAt: []time.Time{request.Deadline},
	})
	if err != nil {
		return domain.CreateGoalResponse{}, err
	}
	if !foundGoals.IsEmpty() {
		return domain.CreateGoalResponse{}, domain.ErrGoalAlreadyExists
	}

	// (3) 새로운 목표 생성
	createdGoal, err := s.goalStore.CreateGoal(ctx, domain.CreateGoalParams{
		GroupID:  request.GroupID,
		Page:     request.Page,
		Deadline: request.Deadline,
	})

	return domain.CreateGoalResponse{
		GoalID:   createdGoal.ID,
		Page:     createdGoal.Page,
		Deadline: createdGoal.Deadline,
	}, nil
}

func (s *GoalService) ListGoals(ctx context.Context, request domain.ListGoalsRequest) (domain.ListGoalsResponse, error) {
	// (1) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{request.GroupID},
	})
	if err != nil {
		return domain.ListGoalsResponse{}, err
	}
	if foundUsers.IsEmpty() {
		return domain.ListGoalsResponse{}, domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.ListGoalsResponse{}, domain.ErrUserNotInGroup
	}

	// (2) 요청한 구룹의 목표 목록 조회
	foundGroups, err := s.groupStore.ListGroups(ctx, domain.ListGroupsParams{
		IDs:   []uint{request.GroupID},
		Limit: 1,

		WithGoals:      true,
		WithGoalsLimit: request.Limit,
	})
	if err != nil {
		return domain.ListGoalsResponse{}, err
	}
	if foundGroups.IsEmpty() {
		return domain.ListGoalsResponse{}, domain.ErrGroupNotFound
	}

	goalsResponse := make([]domain.GoalResponse, 0, len(foundGroups.First().Goals))
	for _, foundGoal := range foundGroups.First().Goals {
		goalsResponse = append(goalsResponse, domain.GoalResponse{
			GoalID:   foundGoal.ID,
			Page:     foundGoal.Page,
			Deadline: foundGoal.Deadline,
		})
	}

	return domain.ListGoalsResponse{Goals: goalsResponse}, nil
}

func (s *GoalService) PatchGoal(ctx context.Context, request domain.PatchGoalRequest) error {
	// (1) 수정하고자 하는 목표가 존재하는지 확인
	foundGoals, err := s.goalStore.ListGoals(ctx, domain.ListGoalsParmas{
		IDs:   []uint{request.GoalID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundGoals.IsEmpty() {
		return domain.ErrGoalNotFound
	}

	// (2) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundGoals.First().GroupID},
	})
	if err != nil {
		return domain.ErrUserNotFound
	}
	if foundUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.ErrUserNotInGroup
	}

	// (3) 목표 내용을 변경
	err = s.goalStore.PatchGoal(ctx, domain.PatchGoalParams{
		ID:       request.GoalID,
		Page:     request.Page,
		Deadline: request.Deadline,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *GoalService) DeleteGoal(ctx context.Context, request domain.DeleteGoalRequest) error {
	// (1) 삭제하고자 하는 목표가 존재하는지 확인
	foundGoals, err := s.goalStore.ListGoals(ctx, domain.ListGoalsParmas{
		IDs:   []uint{request.GoalID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundGoals.IsEmpty() {
		return domain.ErrGoalNotFound
	}

	// (2) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:             []uint{request.RequestUserID},
		Limit:           1,
		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundGoals.First().GroupID},
	})
	if err != nil {
		return err
	}
	if foundUsers.IsEmpty() {
		return domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.ErrUserNotInGroup
	}

	// (3) 목표 삭제
	err = s.goalStore.DeleteGoal(ctx, domain.DeleteGoalParams{
		ID: request.GoalID,
	})
	if err != nil {
		return err
	}

	return nil
}
