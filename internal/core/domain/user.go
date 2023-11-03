package domain

type User struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}

func NewUser(username string, id string) User {
	return User{
		Username: username,
		Id:       id,
	}
}
