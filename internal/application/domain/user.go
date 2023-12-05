package domain

type User struct {
	Id       string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
