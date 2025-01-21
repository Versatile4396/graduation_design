package logic

import (
	mysql "forum/dao"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

const secret = "mande.versatile"

func SignUp(fo *models.User) (userId uint64, err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExit(fo)
	if err != nil {
		// 数据库查询出错
		return
	}
	// 生成Uid
	userId, err = snowflake.GetID()
	if err != nil {
		err = mysql.ErrorGenIDFailed
		return
	}
	fo.UserId = userId
	// 生成加密密码
	encryPassword := mysql.EncryptPassword([]byte(fo.Password))
	fo.Password = encryPassword

	// 插入数据库
	err = global.Db.Create(&fo).Error

	return
}

func Login(p *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	return
}
