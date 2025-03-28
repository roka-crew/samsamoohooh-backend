package domain

import "time"

type CreateGoalRequest struct {
	RequestUserID uint      `json:"-"`
	GroupID       uint      `json:"groupID"`
	Deadline      time.Time `json:"deadline"`
	BookPage      int       `json:"bookPage"`
}

type CreateGoalResponse struct {
	GoalID   uint `json:"goalID"`
	BookPage int  `json:"bookPage"`
	Deadline int  `json:"deadline"`
}

type ListGoalsRequest struct {
	RequestUserID uint `json:"-"`
	GoalIDs       []uint
	Limit         int
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
	GoalID        uint       `json:"goalID"`
	BookPage      *int       `json:"bookPage"`
	Deadline      *time.Time `json:"deadline"`
}
type DeleteGoalRequest struct {
	RequestUserID uint `json:"-"`
	GoalID        uint `json:"goalID"`
}
