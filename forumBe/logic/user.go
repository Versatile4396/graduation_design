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
		UserId:       fo.UserId,
		Avatar:       fo.Avatar,
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
