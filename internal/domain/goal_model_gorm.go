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
	GoalID        = "id"
	GoalPage      = "page"
	GoalDeadline  = "deadline"
	GoalCreatedAt = "created_at"
	GoalUpdatedAt = "updated_at"
	GoalDeletedAt = "deleted_at"
	GoalUserID    = "user_id"
	GoalGroupID   = "group_id"
)
