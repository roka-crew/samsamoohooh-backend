package store

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/postgres"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/samber/lo"
)

type UserStore struct {
	db *postgres.Postgres
}

func NewUserStore(
	db *postgres.Postgres,
) (*UserStore, error) {
	return &UserStore{
		db: db,
	}, nil
}

func (s UserStore) CreateUser(ctx context.Context, params domain.CreateUserParams) (domain.User, error) {
	err := s.db.WithContext(ctx).Create(&params).Error
	if err != nil {

		return domain.User{}, apperr.NewInternalError(err)
	}

	return params, nil
}

func (s UserStore) ListUsers(ctx context.Context, params domain.ListUsersParams) (domain.Users, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.Nicknames) > 0 {
		db = db.Where("nickname IN ?", params.Nicknames)
	}

	if len(params.Biographies) > 0 {
		db = db.Where("biography IN ?", params.Biographies)
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

	var users domain.Users
	if err := db.Find(&users).Error; err != nil {
		return nil, apperr.NewInternalError(err)
	}

	return users, nil
}

func (s UserStore) PatchUser(ctx context.Context, params domain.PatchUserParams) error {
	var updates = make(map[string]any)

	if params.Nickname != nil {
		updates["nickname"] = lo.FromPtr(params.Nickname)
	}

	if params.Biography != nil {
		updates["biography"] = lo.FromPtr(params.Biography)
	}

	if err := s.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", params.ID).Updates(updates).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (u UserStore) DeleteUser(ctx context.Context, params domain.DeleteUserParams) error {
	db := u.db.WithContext(ctx)

	if params.ID > 0 {
		db = db.Where("id = ?", params.ID)
	}

	if params.Nickname != "" {
		db = db.Where("nickname = ?", params.Nickname)
	}

	if params.IsHardDelete {
		db = db.Unscoped()
	}

	if err := db.Delete(&domain.User{}).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}
