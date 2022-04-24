package userservice

import (
	"demo/config/db"
	usermodel "demo/models/user_model"
	"demo/util"
)

type UserService struct{}

func (UserService) GetUsers() ([]usermodel.UserModel, error) {
	table := db.GetTable("users")
	var users []usermodel.UserModel
	res := table.Where("age >= ?", "20").Scan(&users)

	if res.Error != nil {
		util.PrintError(res.Error)
		return nil, res.Error
	}

	return users, nil
}
