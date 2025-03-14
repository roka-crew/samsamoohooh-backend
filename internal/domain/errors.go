package domain

import "github.com/roka-crew/samsamoohooh-backend/pkg/apperr"

var (
	ErrUserDuplicate     = apperr.New("ERR_UESR_DUPLICATE")      // 중복된 사용자가 존재할 때
	ErrUserAlreadyExists = apperr.New("ERR_USER_ALREADY_EXISTS") // 이미 존재하는 사용자가 존재할 때
	ErrUserNotFound      = apperr.New("ERR_USER_NOT_FOUND")      // 존재하지 않은 사용자가 존재할 때
)
