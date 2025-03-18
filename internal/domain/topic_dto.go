package domain

type CreateTopicRequest struct {
	UserID  uint   `json:"-"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
