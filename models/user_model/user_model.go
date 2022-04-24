package usermodel

import "time"

type UserModel struct {
	Id        int        `json:id`
	Name      string     `json:name`
	Age       int        `json:age`
	CreatedAt *time.Time `json:created_at`
}
