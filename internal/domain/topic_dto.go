package domain

type CreateTopicRequest struct {
	UserID  uint   `json:"-"`
	GroupID uint   `json:"groupID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListTopicsRequest struct {
	UserID  uint `json:"-"`
	GroupID uint `json:"groupID"`
	Limit   int  `json:"limit"`
}

type ListTopicsResponse struct {
}

type PatchTopicRequest struct {
}

type DeleteTopicRequest struct {
}
