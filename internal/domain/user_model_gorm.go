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

func (u Users) First() User {
	if u.IsEmpty() {
		return User{}
	}
	return u[0]
}

func (u Users) Last() User {
	if u.IsEmpty() {
		return User{}
	}
	return u[len(u)-1]
}

const (
	ModelUserID        = "id"
	ModelUserNickname  = "nickname"
	ModelUserBiography = "biography"
	ModelUserCreatedAt = "created_at"
	ModelUserUpdatedAt = "updated_at"
	ModelUserDeletedAt = "deleted_at"
)
