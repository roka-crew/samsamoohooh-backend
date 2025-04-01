package domain

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Nicknames() []string {
	nicknames := make([]string, 0, len(u))
	for _, user := range u {
		nicknames = append(nicknames, user.Nickname)
	}

	return nicknames
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

func (g Users) Last() User {
	if g.IsEmpty() {
		return User{}
	}
	return g[len(g)-1]
}

const (
	ModelUserID        = "id"
	ModelUserNickname  = "nickname"
	ModelUserBiography = "biography"
	ModelUserCreatedAt = "created_at"
	ModelUserUpdatedAt = "updated_at"
	ModelUserDeletedAt = "deleted_at"
)
