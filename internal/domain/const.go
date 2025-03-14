package domain

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

func (s SortOrder) ToString() string {
	switch s {
	case SortOrderAsc, SortOrderDesc:
		return string(s)
	}

	return string(SortOrderAsc)
}
