package domain

type Goals []Goal

func (g Goals) Len() int {
	return len(g)
}

func (g Goals) IsEmpty() bool {
	return len(g) == 0
}

func (g Goals) First() Goal {
	if g.IsEmpty() {
		return Goal{}
	}
	return g[0]
}

func (g Goals) Last() Goal {
	if g.IsEmpty() {
		return Goal{}
	}
	return g[len(g)-1]
}

const (
	ModelGoalID        = "id"
	ModelGoalPage      = "page"
	ModelGoalDeadline  = "deadline"
	ModelGoalStatus    = "status"
	ModelGoalCreatedAt = "created_at"
	ModelGoalUpdatedAt = "updated_at"
	ModelGoalDeletedAt = "deleted_at"
	ModelGoalUserID    = "user_id"
	ModelGoalGroupID   = "group_id"
)
