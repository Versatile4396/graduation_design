package logic

import (
	mysql "forum/dao"
	"forum/global"
	"forum/models"
	"forum/pkg/jwt"
	"forum/pkg/snowflake"
)

const secret = "mande.versatile"

func SignUp(fo *models.User) (user models.User, err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExit(fo)
	if err != nil {
		// 数据库查询出错
		return
	}
	// 生成Uid
	userId, err := snowflake.GetID()
	if err != nil {
		err = mysql.ErrorGenIDFailed
		return
	}
	fo.UserId = userId
	// 生成加密密码
	encryPassword := mysql.EncryptPassword([]byte(fo.Password))
	fo.Password = encryPassword
	aToken, rToken, err := jwt.GenToken(fo.UserId, fo.UserName)
	if err != nil {
		return
	}
	fo.AccessToken = aToken
	fo.RefreshToken = rToken
	user = models.User{
		AccessToken:  fo.AccessToken,
		RefreshToken: fo.RefreshToken,
		Email:        fo.Email,
		Gender:       fo.Gender,
		UserName:     fo.UserName,
	}
	// 插入数据库
	err = global.Db.Create(&fo).Error

	return user, nil
}

func Login(u *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: u.UserName,
		Password: u.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//  生产JWT
	accessToken, refreshToken, err := jwt.GenToken(user.UserId, u.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
