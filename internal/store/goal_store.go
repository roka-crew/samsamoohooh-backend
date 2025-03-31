package store

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/postgres"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/samber/lo"
)

type GoalStore struct {
	db *postgres.Postgres
}

func NewGoalStore(
	db *postgres.Postgres,
) *GoalStore {
	return &GoalStore{
		db: db,
	}
}

func (s GoalStore) CreateGoal(ctx context.Context, params domain.CreateGoalParams) (domain.Goal, error) {
	err := s.db.WithContext(ctx).Create(&params).Error
	if err != nil {
		return domain.Goal{}, apperr.NewInternalError(err)
	}

	return params, nil
}

func (s GoalStore) ListGoals(ctx context.Context, params domain.ListGoalsParmas) (domain.Goals, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.Pages) > 0 {
		db = db.Where("page IN ?", params.Pages)
	}

	if len(params.Deadlines) > 0 {
		db = db.Where("deadline IN ?", params.Deadlines)
	}

	if len(params.Statuses) > 0 {
		db = db.Where("status IN ?", params.Statuses)
	}

	if len(params.GroupIDs) > 0 {
		db = db.Where("group_id IN ?", params.GroupIDs)
	}

	if len(params.GtCreatedAt) > 0 {
		db = db.Where("created_at > ?", params.GtCreatedAt)
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

	var goals domain.Goals
	if err := db.Find(&goals).Error; err != nil {
		return domain.Goals{}, apperr.NewInternalError(err)
	}

	return goals, nil
}

func (s GoalStore) PatchGoal(ctx context.Context, params domain.PatchGoalParams) error {
	var updates = make(map[string]any)

	if params.Page != nil {
		updates[domain.GoalPage] = lo.FromPtr(params.Page)
	}

	if params.Deadline != nil {
		updates[domain.GoalDeadline] = lo.FromPtr(params.Deadline)
	}

	if err := s.db.Model(domain.Goal{}).Where("id = ?", params.ID).Updates(updates).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GoalStore) DeleteGoal(ctx context.Context, params domain.DeleteGoalParams) error {
	db := s.db.WithContext(ctx)

	if params.IsHardDelete {
		db = db.Unscoped()
	}

	if err := db.Delete(&domain.Goal{}, params.ID).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}
