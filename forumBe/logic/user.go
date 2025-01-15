package logic

import (
	mysql "forum/dao"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

func SignUp(fo *models.User) (error error) {
	// 判断用户是否存在
	err := mysql.CheckUserExit(fo)
	if err != nil {
		// 数据库查询出错
		return err
	}
	// 生成Uid
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}
	fo.UserId = userId
	return global.Db.Create(&fo).Error
}
