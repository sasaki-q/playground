package userservice

import (
	usermodel "demo/models/user_model"
)

type UserService struct{}

func (UserService) GetSelectedUser() *usermodel.UserModel {
	var selectedUser usermodel.UserModel = usermodel.UserModel{
		Id:   1,
		Name: "selected user",
		Age:  20,
	}
	return &selectedUser
}
