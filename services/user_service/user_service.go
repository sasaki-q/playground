package userservice

import (
	"demo/config/db"
	usermodel "demo/models/user_model"
	"demo/util"
)

type UserService struct{}

func (UserService) GetUsers(users *[]usermodel.UserModel) error {
	table := db.GetTable("users")
	res := table.Where("age >= ?", "20").Scan(&users)

	if res.Error != nil {
		util.PrintError(res.Error)
		return res.Error
	}

	return nil
}
