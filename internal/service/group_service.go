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

	// (2) 새로운 구룹에 만든 사용자 추가
	err = s.groupStore.AppendUser(ctx, domain.AppendUserParams{
		GroupID: createdGroup.ID,
		UserIDs: []uint{request.RequestUserID},
	})
	if err != nil {
		return domain.CreateGroupResponse{}, err
	}

	return domain.CreateGroupResponse{
		GroupID:         createdGroup.ID,
		Introduction:    lo.FromPtr(createdGroup.Introduction),
		BookTitle:       createdGroup.BookTitle,
		BookAuthor:      createdGroup.BookAuthor,
		BookPublisher:   lo.FromPtr(createdGroup.BookPublisher),
		BookMaxPage:     createdGroup.BookMaxPage,
		BookCurrentPage: createdGroup.BookCurrentPage,
	}, nil
}

func (s GroupService) ListGroups(ctx context.Context, request domain.ListGroupsRequest) (domain.ListGroupsResponse, error) {
	// (1) 사용자의 구룹 정보를 조회
	foundUsers, err := s.userStore.ListUsers(ctx, domain.ListUsersParams{
		WithGroups:      true,
		WithGroupsLimit: request.Limit,

		IDs:   []uint{request.RequesterID},
		Limit: 1,
	})
	if err != nil {
		return domain.ListGroupsResponse{}, err
	}
	if foundUsers.IsEmpty() {
		return domain.ListGroupsResponse{}, domain.ErrUserNotFound
	}
	if foundUsers.First().Groups.IsEmpty() {
		return domain.ListGroupsResponse{}, domain.ErrGroupNotFound
	}

	groupsResponse := make([]domain.GroupResponse, 0, len(foundUsers.First().Groups))
	for _, group := range foundUsers.First().Groups {
		groupsResponse = append(groupsResponse, domain.GroupResponse{
			GroupID:         group.ID,
			BookTitle:       group.BookTitle,
			BookAuthor:      group.BookAuthor,
			BookPublisher:   lo.FromPtr(group.BookPublisher),
			BookMaxPage:     group.BookMaxPage,
			BookCurrentPage: group.BookCurrentPage,
			Introduction:    lo.FromPtr(group.Introduction),
		})
	}

	return domain.ListGroupsResponse{
		Groups: groupsResponse,
	}, nil
}

func (s GroupService) PatchGroup(ctx context.Context, request domain.PatchGroupRequest) error {
	// (1) 요청한 사용자가, 변경하고자 하는 구룹에 속해있는지 확인
	foundGroups, err := s.groupStore.ListGroups(ctx, domain.ListGroupsParams{
		WithUsers:    true,
		WithUsersIDs: []uint{request.RequestUserID},

		IDs:   []uint{request.GrouopID},
		Limit: 1,
	})
	if err != nil {
		return err
	}
	if foundGroups.IsEmpty() {
		return domain.ErrGroupNotFound
	}
	if foundGroups.First().Users.IsEmpty() {
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
	// // (1) 요청한 사용자가 이미 참가한 구룹이 있는지 확인
	// fetchedGroups, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
	// 	UserID: request.RequesterID,
	// })
	// if err != nil {
	// 	return err
	// }
	// if !fetchedGroups.IsEmpty() {
	// 	return domain.ErrUserAlreadyInGroup
	// }

	// err = s.userStore.AppendGroups(ctx, domain.AppendGroupsParams{
	// 	UserID:   request.RequesterID,
	// 	GroupIDs: request.GroupIDs,
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s GroupService) LeaveGroup(ctx context.Context, request domain.LeaveGroupRequest) error {
	// (1) 요청한 사용자의 탈퇴 구룹 리스트에 속해 있는지 확인
	// fetchedGroups, err := s.userStore.FetchGroups(ctx, domain.FetchGroupsParams{
	// 	UserID: request.RequesterID,
	// })
	// if err != nil {
	// 	return err
	// }
	// if fetchedGroups.IsEmpty() {
	// 	return domain.ErrUserNotInGroup
	// }
	// if !lo.Some(request.GrouopIDs, fetchedGroups.IDs()) {
	// 	return domain.ErrUserNotInGroup
	// }

	// (2) 사용자 구룹에서 나가기
	// err = s.userStore.RemoveGroups(ctx, domain.RemoveGroupsParams{
	// 	UserID:   request.RequesterID,
	// 	GroupIDs: request.GrouopIDs,
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}
