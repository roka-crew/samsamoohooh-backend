package domain

type CreateGroupRequest struct {
	RequesterID   uint    `json:"-"             validate:"required,gte=1"`
	BookTitle     string  `json:"bookTitle"     validate:"required,min=1,max=255"`
	BookAuthor    string  `json:"bookAuthor"    validate:"required,min=1,max=255"`
	BookPublisher *string `json:"bookPublisher" validate:"max=255"`
	BookMaxPage   int     `json:"bookMaxPage"   validate:"required,gte=1"`
	Introduction  *string `json:"introduction"  validate:"max=255"`
}

type CreateGroupResponse struct {
	GroupID         uint   `json:"groupID"`
	BookTitle       string `json:"bookTitle"`
	BookAuthor      string `json:"bookAuthor"`
	BookPublisher   string `json:"bookPublisher"`
	BookMaxPage     int    `json:"bookMaxPage"`
	BookCurrentPage int    `json:"bookCurrentPage"`
	Introduction    string `json:"introduction"`
}

type ListGroupsRequest struct {
	RequesterID uint `json:"-"      validate:"required,gte=1"`
	Limit       int  `query:"limit" validate:"gte=1,lte=300"`
}

type ListGroupsResponse struct {
	Groups []GroupResponse `json:"groups"`
}

type GroupResponse struct {
	GroupID         uint   `json:"groupID"`
	BookTitle       string `json:"bookTitle"`
	BookAuthor      string `json:"bookAuthor"`
	BookPublisher   string `json:"bookPublisher"`
	BookMaxPage     int    `json:"bookMaxPage"`
	BookCurrentPage int    `json:"bookCurrentPage"`
	Introduction    string `json:"introduction"`
}

type PatchGroupRequest struct {
	RequesterID     uint    `json:"-"               validate:"required,gte=1"`
	GrouopID        uint    `params:"group-id"      validate:"required,gte=1"`
	BookTitle       *string `json:"bookTitle"       validate:"min=1,max=255"`
	BookAuthor      *string `json:"bookAuthor"      validate:"min=1,max=255"`
	BookPublisher   *string `json:"bookPublisher"   validate:"max=255"`
	BookMaxPage     *int    `json:"bookMaxPage"     validate:"gte=1"`
	BookCurrentPage *int    `json:"bookCurrentPage" validate:"gte=1"`
	Introduction    *string `json:"introduction"    validate:"max=255"`
}

type JoinGroupRequest struct {
	RequesterID uint   `json:"-" validate:"required,gte=1"`
	GroupIDs    []uint `json:"groupIDs" validate:"required,dive,gte=1"`
}

type LeaveGroupRequest struct {
	RequesterID uint   `json:"-" validate:"required,gte=1"`
	GrouopIDs   []uint `json:"groupIDs" validate:"required,dive,gte=1"`
}
