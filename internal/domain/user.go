package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		Username: name,
		Email:    email,
	}
}
