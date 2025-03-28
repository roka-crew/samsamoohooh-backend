package domain

type CreateTopicRequest struct {
	RequestUserID uint   `json:"-"       validate:"required,gte=1"`
	GroupID       uint   `json:"groupID" validate:"required,gte=1"`
	Title         string `json:"title"   validate:"required,min=4,max=46"`
	Content       string `json:"content" validate:"required,min=4,max=128"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListTopicsRequest struct {
	RequestUserID uint `json:"-"       validate:"required,gte=1"`
	GroupID       uint `json:"groupID" validate:"required,gte=1"`
	Limit         int  `json:"limit"   validate:"required,gte=1,lte=300"`
}

type ListTopicsResponse struct {
	Topics []TopicResponse `json:"topics"`
}

type TopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PatchTopicRequest struct {
	RequestUserID uint   `json:"-"       validate:"required,gte=1"`
	TopicID       uint   `json:"topicID" validate:"required,gte=1"`
	Title         string `json:"title"   validate:"min=4,max=46"`
	Content       string `json:"content" validate:"min=4,max=128"`
}

type DeleteTopicRequest struct {
	RequestUserID uint `json:"-"       validate:"required,gte=1"`
	TopicID       uint `json:"topicID" validate:"required,gte=1"`
}
