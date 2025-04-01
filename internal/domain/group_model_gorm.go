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
	ModelGroupID              = "id"
	ModelGroupIntroduction    = "introduction"
	ModelGroupBookTitle       = "book_title"
	ModelGroupBookAuthor      = "book_author"
	ModelGroupBookPublisher   = "book_publisher"
	ModelGroupBookMaxPage     = "book_max_page"
	ModelGroupBookCurrentPage = "book_current_page"
	ModelGroupCreatedAt       = "created_at"
	ModelGroupUpdatedAt       = "updated_at"
	ModelGroupDeletedAt       = "deleted_at"
)
