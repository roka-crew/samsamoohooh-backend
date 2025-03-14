package domain

import "github.com/roka-crew/samsamoohooh-backend/pkg/apperr"

var (
	// 사용자 관련 에러
	ErrUserDuplicate     = apperr.New("ERR_USER_DUPLICATE")      // 중복된 사용자가 존재할 때
	ErrUserAlreadyExists = apperr.New("ERR_USER_ALREADY_EXISTS") // 이미 존재하는 사용자가 존재할 때
	ErrUserNotFound      = apperr.New("ERR_USER_NOT_FOUND")      // 존재하지 않은 사용자가 존재할 때

	// 인증 관련 에러
	ErrAuthRequired      = apperr.New("ERR_AUTH_REQUIRED")       // 인증 헤더가 없을 때
	ErrAuthInvalidFormat = apperr.New("ERR_AUTH_INVALID_FORMAT") // 잘못된 인증 형식일 때
	ErrAuthInvalidToken  = apperr.New("ERR_AUTH_INVALID_TOKEN")  // 유효하지 않은 토큰일 때
	ErrAuthExpiredToken  = apperr.New("ERR_AUTH_EXPIRED_TOKEN")  // 만료된 토큰일 때
)
