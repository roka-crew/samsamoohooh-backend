package domain

type Topics []Topic

func (t Topics) Len() int {
	return len(t)
}

func (t Topics) IsEmpty() bool {
	return len(t) == 0
}

func (t Topics) First() Topic {
	if t.IsEmpty() {
		return Topic{}
	}
	return t[0]
}

func (t Topics) Last() Topic {
	if t.IsEmpty() {
		return Topic{}
	}
	return t[len(t)-1]
}
