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
