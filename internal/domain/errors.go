package domain

import (
	"net/http"

	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

var (
	// 사용자 관련 에러
	ErrUserDuplicate      = apperr.New("ERR_USER_DUPLICATE").WithStatus(http.StatusConflict)        // 중복된 사용자가 존재할 때
	ErrUserAlreadyExists  = apperr.New("ERR_USER_ALREADY_EXISTS").WithStatus(http.StatusConflict)   // 이미 존재하는 사용자가 존재할 때
	ErrUserNotFound       = apperr.New("ERR_USER_NOT_FOUND").WithStatus(http.StatusNotFound)        // 존재하지 않은 사용자가 존재할 때
	ErrUserNotInGroup     = apperr.New("ERR_USER_NOT_IN_GROUP").WithStatus(http.StatusForbidden)    // 사용자가 해당 그룹에 속해있지 않을 때
	ErrUserAlreadyInGroup = apperr.New("ERR_USER_ALREADY_IN_GROUP").WithStatus(http.StatusConflict) // 사용자가 이미 해당 그룹에 속해있을 때

	// 구룹 관련 에러
	ErrGroupNotFound = apperr.New("ERR_GROUP_NOT_FOUND").WithStatus(http.StatusNotFound) // 구룹이 존재하지 않을 때

	// 목표 관련 에러
	ErrGoalAlreadyExists = apperr.New("ERR_GOAL_ALREADY_EXISTS").WithStatus(http.StatusConflict) // 이미 존재하는 목표가 있을 때
	ErrGoalNotFound      = apperr.New("ERR_GOAL_NOT_FOUND").WithStatus(http.StatusNotFound)      // 목표가 존재하지 않을 때

	// 주제 관련 에러
	ErrTopicNotFound = apperr.New("ERR_TOPIC_NOT_FOUND").WithStatus(http.StatusNotFound) // 주제가 존재하지 않을 때

	// 인증 관련 에러
	ErrAuthRequired             = apperr.New("ERR_AUTH_REQUIRED").WithStatus(http.StatusUnauthorized)               // 인증 헤더가 없을 때
	ErrAuthInvalidFormat        = apperr.New("ERR_AUTH_INVALID_FORMAT").WithStatus(http.StatusBadRequest)           // 잘못된 인증 형식일 때
	ErrAuthInvalidToken         = apperr.New("ERR_AUTH_INVALID_TOKEN").WithStatus(http.StatusUnauthorized)          // 유효하지 않은 토큰일 때
	ErrAuthExpiredToken         = apperr.New("ERR_AUTH_EXPIRED_TOKEN").WithStatus(http.StatusUnauthorized)          // 만료된 토큰일 때
	ErrAuthInvalidSigningMethod = apperr.New("ERR_AUTH_INVALID_SIGNING_METHOD").WithStatus(http.StatusUnauthorized) // 잘못된 서명 방식일 때
	ErrAuthInvalidClaims        = apperr.New("ERR_AUTH_INVALID_CLAIMS").WithStatus(http.StatusUnauthorized)         // 토큰 클레임이 잘못되었을 때
	ErrAuthMalformedToken       = apperr.New("ERR_AUTH_MALFORMED_TOKEN").WithStatus(http.StatusBadRequest)          // 토큰 형식이 잘못되었을 때
)
