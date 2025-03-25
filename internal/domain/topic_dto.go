package domain

type CreateTopicRequest struct {
	RequesterID uint   `json:"-"       validate:"required,gte=1"`
	GroupID     uint   `json:"groupID" validate:"required,gte=1"`
	Title       string `json:"title"   validate:"required,min=4,max=46"`
	Content     string `json:"content" validate:"required,min=4,max=128"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListTopicsRequest struct {
	RequesterID uint `json:"-" validate:"required,gte=1"`
	GroupID     uint `json:"groupID" validate:"required,gte=1"`
	Limit       int  `json:"limit"  validate:"required,gte=1,lte=300"`
}

type ListTopicsResponse struct {
}

type PatchTopicRequest struct {
}

type DeleteTopicRequest struct {
}
