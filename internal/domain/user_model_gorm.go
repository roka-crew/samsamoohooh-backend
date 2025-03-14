package domain

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) IsEmpty() bool {
	return len(u) == 0
}

func (t Users) First() User {
	if t.IsEmpty() {
		return User{}
	}
	return t[0]
}

func (t Users) Last() User {
	if t.IsEmpty() {
		return User{}
	}
	return t[len(t)-1]
}
