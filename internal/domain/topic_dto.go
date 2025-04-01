package domain

type CreateTopicRequest struct {
	RequestUserID uint   `json:"-"       validate:"required,gte=1"`
	GoalID        uint   `json:"goalID"  validate:"required,gte=1"`
	Title         string `json:"title"   validate:"required,min=4,max=46"`
	Content       string `json:"content" validate:"required,min=4,max=128"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topicID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListTopicsRequest struct {
	RequestUserID uint `json:"-"       validate:"required,gte=1"`
	GoalID        uint `query:"goalID" validate:"required,gte=1"`
	Limit         int  `query:"limit"  validate:"required,gte=1,lte=300"`
}

type ListTopicsResponse struct {
	Topics []TopicResponse `json:"topics"`
}

type TopicResponse struct {
	TopicID uint   `json:"topicID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PatchTopicRequest struct {
	RequestUserID uint    `json:"-"         validate:"required,gte=1"`
	TopicID       uint    `params:"topicID" validate:"required,gte=1"`
	Title         *string `json:"title"     validate:"omitempty,min=4,max=46"`
	Content       *string `json:"content"   validate:"omitempty,min=4,max=128"`
}

type DeleteTopicRequest struct {
	RequestUserID uint `json:"-"       validate:"required,gte=1"`
	TopicID       uint `json:"topicID" validate:"required,gte=1"`
}
