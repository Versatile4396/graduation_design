package logic

import (
	mysql "forum/dao"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

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
	err = global.Db.Create(&fo).Error
	return
}
