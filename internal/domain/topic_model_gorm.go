package domain

type Topics []Topic

func (t Topics) Titles() []string {
	titles := make([]string, 0, len(t))
	for _, topic := range t {
		titles = append(titles, topic.Title)
	}

	return titles
}

func (t Topics) Len() int {
	return len(t)
}

func (t Topics) IsEmpty() bool {
	return len(t) == 0
}

func (t Topics) First() Topic {
	if t.IsEmpty() {
		return Topic{}
	}
	return t[0]
}

func (t Topics) Last() Topic {
	if t.IsEmpty() {
		return Topic{}
	}
	return t[len(t)-1]
}

const (
	ModelTopicID        = "id"
	ModelTopicTitle     = "title"
	ModelTopicContent   = "content"
	ModelTopicCreatedAt = "created_at"
	ModelTopicUpdatedAt = "updated_at"
	ModelTopicDeletedAt = "deleted_at"

	ModelTopicGoalID = "goal_id"
	ModelTopicUserID = "user_id"
)
