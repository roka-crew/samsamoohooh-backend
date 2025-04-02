package service

import (
	"context"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
)

type TopicService struct {
	topicStore *store.TopicStore
	groupStore *store.GroupStore
	userStore  *store.UserStore
	goalStore  *store.GoalStore
}

func NewTopicService(
	topicStore *store.TopicStore,
	groupStore *store.GroupStore,
	userStore *store.UserStore,
	goalStore *store.GoalStore,
) *TopicService {
	return &TopicService{
		topicStore: topicStore,
		groupStore: groupStore,
		userStore:  userStore,
		goalStore:  goalStore,
	}
}

func (s TopicService) CreateTopic(ctx context.Context, request domain.CreateTopicRequest) (domain.CreateTopicResponse, error) {
	// (1) 요청한 사용자가 생성할 목표의 구룹에 속해있는지 확인
	foundGoals, err := s.goalStore.ListGoals(ctx, domain.ListGoalsParams{
		IDs:   []uint{request.GoalID},
		Limit: 1,
	})
	if err != nil {
		return domain.CreateTopicResponse{}, err
	}
	if foundGoals.IsEmpty() {
		return domain.CreateTopicResponse{}, domain.ErrGoalNotFound
	}
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundGoals.First().GroupID},
	})
	if err != nil {
		return domain.CreateTopicResponse{}, err
	}
	if foundUsers.IsEmpty() {
		return domain.CreateTopicResponse{}, domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.CreateTopicResponse{}, domain.ErrUserNotInGroup
	}

	// (2) 토픽 생성
	createdTopic, err := s.topicStore.CreateTopic(ctx, domain.CreateTopicParams{
		GoalID:  request.GoalID,
		UserID:  request.RequestUserID,
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		return domain.CreateTopicResponse{}, err
	}

	return domain.CreateTopicResponse{
		TopicID: createdTopic.ID,
		Title:   createdTopic.Title,
		Content: createdTopic.Content,
	}, nil
}

func (s TopicService) ListTopics(ctx context.Context, request domain.ListTopicsRequest) (domain.ListTopicsResponse, error) {
	// (1) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundGoals, err := s.goalStore.ListGoals(ctx, domain.ListGoalsParams{
		IDs:   []uint{request.GoalID},
		Limit: 1,
	})
	if err != nil {
		return domain.ListTopicsResponse{}, err
	}
	if foundGoals.IsEmpty() {
		return domain.ListTopicsResponse{}, domain.ErrGoalNotFound
	}
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundGoals.First().GroupID},
	})
	if err != nil {
		return domain.ListTopicsResponse{}, err
	}
	if foundUsers.IsEmpty() {
		return domain.ListTopicsResponse{}, domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.ListTopicsResponse{}, domain.ErrUserNotInGroup
	}

	// (2) 요청한 목표에 대한 토픽 목록을 조회
	foundTopics, err := s.topicStore.ListTopics(ctx, domain.ListTopicsParams{
		GoalIDs: []uint{request.GoalID},
		Limit:   request.Limit,

		Order:   domain.SortOrderDesc,
		OrderBy: domain.ModelTopicCreatedAt,
	})
	if err != nil {
		return domain.ListTopicsResponse{}, err
	}

	// (3) 요청한 목표에 대한 토픽 목록을 응답으로 반환
	topics := make([]domain.TopicResponse, 0, len(foundTopics))
	for _, foundTopic := range foundTopics {
		topics = append(topics, domain.TopicResponse{
			TopicID: foundTopic.ID,
			Title:   foundTopic.Title,
			Content: foundTopic.Content,
		})
	}

	return domain.ListTopicsResponse{Topics: topics}, nil
}

func (s TopicService) PatchTopic(ctx context.Context, request domain.PatchTopicRequest) error {
	// (1) 수정하고자 하는 목표가 존재하는지 확인
	foundTopics, err := s.topicStore.ListTopics(ctx, domain.ListTopicsParams{
		IDs:   []uint{request.TopicID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundTopics.IsEmpty() {
		return domain.ErrTopicNotFound
	}

	// (2) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundTopics.First().GoalID},
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

	// (3) 토픽 수정
	err = s.topicStore.PatchTopic(ctx, domain.PatchTopic{
		ID:      request.TopicID,
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s TopicService) DeleteTopic(ctx context.Context, request domain.DeleteTopicRequest) error {
	// (1) 수정하고자 하는 목표가 존재하는지 확인
	foundTopics, err := s.topicStore.ListTopics(ctx, domain.ListTopicsParams{
		IDs:   []uint{request.TopicID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundTopics.IsEmpty() {
		return domain.ErrTopicNotFound
	}

	// (2) 요청한 사용자가 해당 구룹에 속해있는지 확인
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		IDs:   []uint{request.RequestUserID},
		Limit: 1,

		WithGroups:      true,
		WithGroupsLimit: 1,
		WithGroupsIDs:   []uint{foundTopics.First().GoalID},
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

	// (3) 토픽 삭제
	err = s.topicStore.DeleteTopic(ctx, domain.DeleteTopic{
		ID: request.TopicID,
	})
	if err != nil {
		return err
	}

	return nil
}
