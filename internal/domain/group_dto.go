package domain

type CreateGroupRequest struct {
	UserID        uint    `json:"-"`
	BookTitle     string  `json:"bookTitle" validate:"required"`
	BookAuthor    string  `json:"bookAuthor" validate:"required"`
	BookPublisher *string `json:"bookPublisher"`
	BookMaxPage   int     `json:"bookMaxPage" validate:"required"`
	Introduction  *string `json:"introduction"`
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
	UserID uint `json:""`
	Limit  int  `query:"limit"`
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
	UserID          uint    `json:"-"`
	GrouopID        uint    `params:"group-id"`
	BookTitle       *string `json:"bookTitle"`
	BookAuthor      *string `json:"bookAuthor"`
	BookPublisher   *string `json:"bookPublisher"`
	BookMaxPage     *int    `json:"bookMaxPage"`
	BookCurrentPage *int    `json:"bookCurrentPage"`
	Introduction    *string `json:"introduction"`
}

type JoinGroupRequest struct {
	UserID   uint   `json:"-"`
	GroupIDs []uint `json:"groupIDs"`
}

type LeaveGroupRequest struct {
	UserID    uint   `json:"-"`
	GrouopIDs []uint `json:"groupIDs"`
}
