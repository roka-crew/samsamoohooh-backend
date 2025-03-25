package store

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/postgres"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/samber/lo"
)

type GroupStore struct {
	db *postgres.Postgres
}

func NewGroupStore(
	db *postgres.Postgres,
) *GroupStore {
	return &GroupStore{
		db: db,
	}
}

func (s GroupStore) CreateGroup(ctx context.Context, params domain.CreateGroupParams) (domain.Group, error) {
	err := s.db.WithContext(ctx).Create(&params).Error
	if err != nil {
		return domain.Group{}, apperr.NewInternalError(err)
	}

	return params, nil
}

func (s GroupStore) ListGroups(ctx context.Context, params domain.ListGroupsParams) (domain.Groups, error) {
	db := s.db.WithContext(ctx)

	if len(params.IDs) > 0 {
		db = db.Where("id IN ?", params.IDs)
	}

	if len(params.BookTitles) > 0 {
		db = db.Where("book_title IN ?", params.BookTitles)
	}

	if len(params.BookAuthors) > 0 {
		db = db.Where("book_author IN ?", params.BookAuthors)
	}

	if len(params.BookPublishers) > 0 {
		db = db.Where("book_publisher IN ?", params.BookPublishers)
	}

	if len(params.BookMaxPages) > 0 {
		db = db.Where("book_max_page IN ?", params.BookMaxPages)
	}

	if len(params.BookCurrentPages) > 0 {
		db = db.Where("book_current_page IN ?", params.BookCurrentPages)
	}

	if params.WithGoals {
		db = db.Preload("Goals")
	}

	if params.WithUsers {
		db = db.Preload("Users")
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

	var groups domain.Groups
	if err := db.Find(&groups).Error; err != nil {
		return domain.Groups{}, apperr.NewInternalError(err)
	}

	return groups, nil
}

func (s GroupStore) PatchGroup(ctx context.Context, params domain.PatchGroupParams) error {
	var updates = make(map[string]any)

	if params.Introduction != nil {
		updates["introduction"] = lo.FromPtr(params.Introduction)
	}

	if params.BookTitle != nil {
		updates["book_title"] = lo.FromPtr(params.BookTitle)
	}

	if params.BookAuthor != nil {
		updates["book_author"] = lo.FromPtr(params.BookAuthor)
	}

	if params.BookPublisher != nil {
		updates["book_publisher"] = lo.FromPtr(params.BookPublisher)
	}

	if params.BookCurrentPage != nil {
		updates["book_current_page"] = lo.FromPtr(params.BookCurrentPage)
	}

	if err := s.db.WithContext(ctx).Updates(updates).Error; err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) DeleteGroup(ctx context.Context, params domain.DeleteGroupParams) error {
	db := s.db.WithContext(ctx)

	if params.IsHardDelete {
		db = db.Unscoped()
	}

	err := db.Delete(&domain.User{}).Where("id = ?", params.ID).Error
	if err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) AppendUser(ctx context.Context, params domain.AppendUserParams) error {
	wantAppendUser := domain.Users{}
	for _, userID := range params.UserIDs {
		wantAppendUser = append(wantAppendUser, domain.User{
			ID: userID,
		})
	}

	err := s.db.WithContext(ctx).
		Model(&domain.Group{ID: params.GroupID}).
		Association("Users").
		Append(&wantAppendUser)
	if err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) RemoveUsers(ctx context.Context, params domain.RemoveUsersParams) error {
	wantRemoveUsers := domain.Users{}
	for _, userID := range params.UserIDs {
		wantRemoveUsers = append(wantRemoveUsers, domain.User{
			ID: userID,
		})
	}

	err := s.db.WithContext(ctx).
		Model(&domain.Group{ID: params.GroupID}).
		Association("Groups").
		Delete(wantRemoveUsers)
	if err != nil {
		return apperr.NewInternalError(err)
	}

	return nil
}

func (s GroupStore) FetchUsers(ctx context.Context, params domain.FetchUsersParams) (domain.Users, error) {
	db := s.db.WithContext(ctx)

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	var users domain.Users
	err := db.
		Model(&domain.Group{}).
		Where("id IN ?", params.GroupIDs).
		Association("Users").
		Find(&users)
	if err != nil {
		return domain.Users{}, apperr.NewInternalError(err)
	}

	return users, nil
}
