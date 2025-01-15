package mysql

import (
	"errors"
	"forum/global"
	"forum/models"
)

func CheckUserExit(user *models.User) (error error) {
	var count int64
	result := global.Db.Model(&models.User{}).Where("user_name =?", user.UserName).Count(&count)
	if result.Error != nil {
		panic(result.Error)
	}
	if count > 0 {
		return errors.New(ErrorUserExit)
	}
	return
}
