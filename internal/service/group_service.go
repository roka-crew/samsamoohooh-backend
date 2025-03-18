package service

import (
	"context"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/store"
	"github.com/samber/lo"
)

type GroupService struct {
	userStore  *store.UserStore
	groupStore *store.GroupStore
}

func NewGroupService(
	userStore *store.UserStore,
	groupStore *store.GroupStore,
) *GroupService {
	return &GroupService{
		userStore:  userStore,
		groupStore: groupStore,
	}
}

func (s GroupService) CreateGroup(ctx context.Context, request domain.CreateGroupRequest) (domain.CreateGroupResponse, error) {
	// (1) 새로운 구룹을 생성
	createdGroup, err := s.groupStore.CreateGroup(ctx, domain.CreateGroupParams{
		Introduction:    request.Introduction,
		BookTitle:       request.BookTitle,
		BookAuthor:      request.BookAuthor,
		BookPublisher:   request.BookPublisher,
		BookMaxPage:     request.BookMaxPage,
		BookCurrentPage: 0,
	})
	if err != nil {
		return domain.CreateGroupResponse{}, err
	}

	// (2) 새로운 구룹에 사용자 추가하기
	err = s.groupStore.AppendUser(ctx, domain.AppendUserParams{
		GroupID: createdGroup.ID,
		UserIDs: []uint{request.UserID},
	})
	if err != nil {
		return domain.CreateGroupResponse{}, err
	}

	return createdGroup.ToCreateGroupResponse(), nil
}

func (s GroupService) ListGroups(ctx context.Context, request domain.ListGroupsRequest) (domain.ListGroupsResponse, error) {
	// (1) 사용자의 구룹 정보를 조회
	fetchedGroups, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
		UserIDs: []uint{request.UserID},
		Limit:   request.Limit,
	})
	if err != nil {
		return domain.ListGroupsResponse{}, err
	}

	return fetchedGroups.ToListGroupsResponse(), nil
}

func (s GroupService) PatchGroup(ctx context.Context, request domain.PatchGroupRequest) error {
	// (1) 요청한 사용자가, 변경하고자 하는 구룹에 속해있는지 확인
	fetchedGrouops, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
		UserIDs: []uint{request.UserID},
		Limit:   1,
	})
	if err != nil {
		return err
	}
	if fetchedGrouops.IsEmpty() {
		return domain.ErrUserNotInGroup
	}

	// (2) 구룹 내용을 변경
	err = s.groupStore.PatchGroup(ctx, domain.PatchGroupParams{
		ID:              request.GrouopID,
		Introduction:    request.Introduction,
		BookTitle:       request.BookTitle,
		BookAuthor:      request.BookAuthor,
		BookPublisher:   request.BookPublisher,
		BookMaxPage:     request.BookMaxPage,
		BookCurrentPage: request.BookCurrentPage,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s GroupService) JoinGroup(ctx context.Context, request domain.JoinGroupRequest) error {
	// (1) 요청한 사용자가 이미 참가한 구룹이 있는지 확인
	fetchedGroups, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
		UserIDs: []uint{request.UserID},
	})
	if err != nil {
		return err
	}
	if !fetchedGroups.IsEmpty() {
		return domain.ErrUserAlreadyInGroup
	}

	// (2) 요청한 사용자를 구룹에 속하게 하기
	err = s.userStore.AppendGroups(ctx, domain.AppendGroupsParams{
		UserID:   request.UserID,
		GroupIDs: request.GroupIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s GroupService) LeaveGroup(ctx context.Context, request domain.LeaveGroupRequest) error {
	// (1) 요청한 사용자의 탈퇴 구룹 리스트에 속해 있는지 확인
	fetchedGroups, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
		UserIDs: []uint{request.UserID},
	})
	if err != nil {
		return err
	}
	if fetchedGroups.IsEmpty() {
		return domain.ErrUserNotInGroup
	}
	if !lo.Some(request.GrouopIDs, fetchedGroups.IDs()) {
		return domain.ErrUserNotInGroup
	}

	// (2) 사용자 구룹에서 나가기
	err = s.userStore.RemoveGroups(ctx, domain.RemoveGroupsParams{
		UserID:   request.UserID,
		GroupIDs: request.GrouopIDs,
	})
	if err != nil {
		return err
	}

	return nil
}
