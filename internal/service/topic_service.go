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
}

func NewTopicService(
	topicStore *store.TopicStore,
	groupStore *store.GroupStore,
	userStore *store.UserStore,
) *TopicService {
	return &TopicService{
		topicStore: topicStore,
		groupStore: groupStore,
		userStore:  userStore,
	}
}

func (s TopicService) CreateTopic(ctx context.Context, request domain.CreateTopicRequest) (domain.CreateTopicResponse, error) {
	// hasGroup, err := s.userStore.HasGroups(ctx, domain.HasGroupsParams{
	// 	UserID:   request.RequestUserID,
	// 	GroupIDs: []uint{request.GroupID},
	// })
	// if err != nil {
	// 	return domain.CreateTopicResponse{}, err
	// }
	// if !hasGroup {
	// 	return domain.CreateTopicResponse{}, domain.ErrUserNotInGroup
	// }

	// return domain.CreateTopicResponse{}, nil
	// (1) 새로운 토픽 생성
	// createdTopic, err := s.topicStore.CreateTopic(ctx, domain.CreateTopicParams{

	// })
	// if err != nil {
	// 	return domain.CreateTopicResponse{}, err
	// }

	// return createdTopic.ToCreateTopicResponse(), nil
	return domain.CreateTopicResponse{}, nil
}
