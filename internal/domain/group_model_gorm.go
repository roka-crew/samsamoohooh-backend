package domain

type Groups []Group

func (g Groups) Len() int {
	return len(g)
}

func (g Groups) IsEmpty() bool {
	return len(g) == 0
}

func (g Groups) First() Group {
	if g.IsEmpty() {
		return Group{}
	}
	return g[0]
}

func (g Groups) Last() Group {
	if g.IsEmpty() {
		return Group{}
	}
	return g[len(g)-1]
}

func (g Groups) IDs() []uint {
	ids := make([]uint, 0, len(g))

	for _, group := range g {
		ids = append(ids, group.ID)
	}

	return ids
}

const (
	GroupID              = "id"
	GroupIntroduction    = "introduction"
	GroupBookTitle       = "book_title"
	GroupBookAuthor      = "book_author"
	GroupBookPublisher   = "book_publisher"
	GroupBookMaxPage     = "book_max_page"
	GroupBookCurrentPage = "book_current_page"
	GroupCreatedAt       = "created_at"
	GroupUpdatedAt       = "updated_at"
	GroupDeletedAt       = "deleted_at"
)
