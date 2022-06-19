package models

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
