package store

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/postgres"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/samber/lo"
)

type TopicStore struct {
	db *postgres.Postgres
}

func NewTopicStore(
	db *postgres.Postgres,
) *TopicStore {
	return &TopicStore{
		db: db,
	}
}

func (s TopicStore) CreateTopic(ctx context.Context, params domain.CreateTopicParams) (domain.Topic, error) {
	err := s.db.WithContext(ctx).Create(&params).Error
	if err != nil {
		return domain.Topic{}, apperr.NewInternalError(err)
	}

	return params, nil
}

func (s TopicStore) ListTopics(ctx context.Context, params domain.ListTopicsParams) (domain.Topics, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.Titles) > 0 {
		db = db.Where("title IN ?", params.Titles)
	}

	if len(params.Contents) > 0 {
		db = db.Where("content IN ?", params.Contents)
	}

	if len(params.GoalIDs) > 0 {
		db = db.Where("goal_id IN ?", params.GoalIDs)
	}

	if params.OrderBy != "" {
		db = db.Order(params.OrderBy + " " + params.Order.ToString())
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	if params.Offset > 0 {
		db = db.Offset(params.Offset)
	}

	var topics domain.Topics
	if err := db.Find(&topics).Error; err != nil {
		return domain.Topics{}, apperr.NewInternalError(err)
	}

	return topics, nil
}

func (s TopicStore) PatchTopic(ctx context.Context, params domain.PatchTopic) error {
	var updates = make(map[string]any)

	if params.Title != nil {
		updates[domain.ModelTopicTitle] = lo.FromPtr(params.Title)
	}

	if params.Content != nil {
		updates[domain.ModelTopicContent] = lo.FromPtr(params.Content)
	}

	err := s.db.WithContext(ctx).Model(&domain.Topic{ID: params.ID}).Updates(updates).Error
	if err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s TopicStore) DeleteTopic(ctx context.Context, params domain.DeleteTopic) error {
	db := s.db.WithContext(ctx)

	if params.IsHardDelete {
		db = db.Unscoped()
	}

	if err := db.Delete(&domain.Topic{ID: params.ID}).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}
