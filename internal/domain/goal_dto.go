package domain

import "time"

type CreateGoalRequest struct {
	RequestUserID uint      `json:"-"        validate:"required,gte=1"`
	GroupID       uint      `json:"groupID"  validate:"required,gte=1"`
	Deadline      time.Time `json:"deadline" validate:"required"`
	BookPage      int       `json:"bookPage" validate:"required,gte=1"`
}

type CreateGoalResponse struct {
	GoalID   uint `json:"goalID"`
	BookPage int  `json:"bookPage"`
	Deadline int  `json:"deadline"`
}

type ListGoalsRequest struct {
	RequestUserID uint   `json:"-"`
	GoalIDs       []uint `json:"goalIDs" validate:"dive,gte=1"`
	Limit         int    `query:"limit"  validate:"gte=1,lte=300"`
}

type ListGoalsResponse struct {
	Goals []GoalResponse `json:"goals"`
}

type GoalResponse struct {
	GoalID   uint `json:"goalID"`
	BookPage int  `json:"bookPage"`
	Deadline int  `json:"deadline"`
}

type PatchGoalRequest struct {
	RequestUserID uint       `json:"-"`
	GoalID        uint       `json:"goalID"   validate:"required,gte=1"`
	BookPage      *int       `json:"bookPage" validate:"gte=1"`
	Deadline      *time.Time `json:"deadline"`
}
type DeleteGoalRequest struct {
	RequestUserID uint `json:"-"`
	GoalID        uint `json:"goalID"`
}
