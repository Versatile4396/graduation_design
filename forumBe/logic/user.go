package logic

import (
	mysql "forum/dao"
	"forum/global"
	"forum/models"
	"forum/pkg/jwt"
	"forum/pkg/snowflake"
	"log"
	"time"
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
		UserId:       fo.UserId,
		Role:         fo.Role,
		Avatar:       fo.Avatar,
		Nickname:     fo.Nickname,
		Overview:     fo.Overview,
	}
	// 插入数据库
	err = global.Db.Create(&fo).Error
	if err != nil {
		return
	}
	return user, nil
}

func Login(u *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: u.UserName,
		Password: u.Password,
	}
	dbuser, err := mysql.Login(user)

	if err != nil {
		return nil, err
	}
	//  生产JWT
	accessToken, refreshToken, err := jwt.GenToken(user.UserId, u.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.UserId = dbuser.UserId
	user.Email = dbuser.Email
	user.UserName = dbuser.UserName
	user.Avatar = dbuser.Avatar
	user.Gender = dbuser.Gender
	user.Nickname = dbuser.Nickname
	user.Overview = dbuser.Overview
	user.Role = dbuser.Role
	return
}

/**
 * @description: 获取用户信息
 * @param c
 * @return {*}
 */
func GetCountController(uid uint64) (countInfo *models.UserCountInfo, err error) {

	countInfo = &models.UserCountInfo{}
	// 获取文章数量 和总浏览量
	rArticle := global.Db.Model(&models.Article{}).Where("user_id =?", uid)
	err = rArticle.Error
	if err != nil {
		return
	}
	rArticle.Count(&countInfo.AritcleCount)
	rArticle.Select("sum(view_count)").Row().Scan(&countInfo.ViewCount)
	// 获取点赞数量
	err = global.Db.Model(&models.ArticleLike{}).
		Joins("JOIN articles ON article_likes.article_id = articles.article_id").
		Where("articles.user_id =?", uid).
		Count(&countInfo.LikeCount).Error
	if err != nil {
		return
	}
	// comments.Count(&countInfo.CommentCount)
	return countInfo, nil
}

func GetFriendList(uid uint64) (userList []models.UserInfo, err error) {
	// 获取用户的好友列表
	var friends []models.UserFriend
	err = global.Db.Model(&models.UserFriend{}).Where("user_id =?", uid).Find(&friends).Error
	for _, friend := range friends {
		var userInfo models.UserInfo
		err = global.Db.Model(&models.User{}).Where("user_id =?", friend.FriendId).First(&userInfo).Error
		userList = append(userList, userInfo)
	}
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func GetUserList(fo *models.UserFilter) (userList []models.UserInfo, err error) {
	if fo.Pagination == nil {
		fo.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := global.Db.Model(&models.User{})
	if fo.UserId != 0 {
		query = query.Where("user_id = ?", fo.UserId)
	}
	if fo.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+fo.UserName+"%")
	}
	if fo.Nickname != "" {
		query = query.Where("nickname LIKE?", "%"+fo.Nickname+"%")
	}
	if fo.Overview != "" {
		query = query.Where("overview LIKE?", "%"+fo.Overview+"%")
	}
	if fo.Email != "" {
		query = query.Where("email LIKE?", "%"+fo.Email+"%")
	}
	if fo.Gender != 0 {
		query = query.Where("gender =?", fo.Gender)
	}
	offset := fo.Pagination.PageSize * (fo.Pagination.Page - 1)
	query.Find(&userList).Offset(offset).Limit(fo.Pagination.PageSize)
	return
}

func UpdateUser(fo *models.UserUpdateForm) (err error) {
	// 更新用户信息
	result := global.Db.Model(&models.User{}).Where("user_id =?", fo.UserId).Updates(fo)
	if result.Error != nil {
		log.Printf("更新用户信息时出错: %v", result.Error)
		return result.Error
	}
	return
}

func DeleteUser(uid uint64, delete int) (err error) {
	// 处理用户信息
	err = global.Db.Model(&models.User{}).Where("user_id =?", uid).Update("deleted", delete).Error
	curTime := time.Now()
	err = global.Db.Model(&models.User{}).Where("user_id =?", uid).Update("deleted_at", curTime).Error
	return
}

func Friend(fo *models.UserFriendForm) (err error) {
	// 前置处理
	err = mysql.CheckFriendsExit(fo)
	if err != nil {
		return err
	}
	userFriend := models.UserFriend{
		UserId:   fo.UserId,
		FriendId: fo.FriendId,
	}
	userFriendReverse := models.UserFriend{
		UserId:   fo.FriendId,
		FriendId: fo.UserId,
	}
	global.Db.Create(&userFriend)
	global.Db.Create(&userFriendReverse)
	return err
}
