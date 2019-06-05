package models

type User struct {
	Id       int64  `json:"id"`
	Password string `password`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
