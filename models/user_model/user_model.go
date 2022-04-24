package usermodel

import "time"

type UserModel struct {
	Id        *int       `json:"id, omitempty"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"created_at, omitempty"`
}

type CreateModel struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
