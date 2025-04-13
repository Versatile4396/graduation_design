package controller

import (
	"forum/global"
	"forum/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowCreateController(c *gin.Context) {
	var follow models.Follow
	if err := c.ShouldBindJSON(&follow); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 检查是否已经关注
	if IsFollowerLogic(follow.FollowerID, follow.FollowedID) {
		ResponseError(c, CodeFollowAlreadyExist)
		return
	}
	if err := global.Db.Create(&follow).Error; err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func FollowDeleteController(c *gin.Context) {
	var follow models.Follow
	if err := c.ShouldBindJSON(&follow); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if err := global.Db.Model(&models.Follow{}).
		Where("follower_id =? AND followed_id=?", follow.FollowerID, follow.FollowedID).
		Delete(&follow).Error; err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func FollowGetListController(c *gin.Context) {
	var uid = c.Param("uid")
	var postId, err = strconv.ParseUint(uid, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 查询关注的用户id列表
	var followList []models.Follow
	var followerUserInfo []models.UserInfo
	if err := global.Db.Where("follower_id =?", postId).Find(&followList).Error; err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	for _, follow := range followList {
		var userInfo models.UserInfo
		if err := global.Db.Model(&models.User{}).Where("user_id =?", follow.FollowedID).First(&userInfo).Error; err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
		followerUserInfo = append(followerUserInfo, userInfo)
	}
	// 查询关注你的用户id列表
	var followedList []models.Follow
	var followedUserInfo []models.UserInfo
	if err := global.Db.Where("followed_id =?", postId).Find(&followedList).Error; err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	for _, follow := range followedList {
		var userInfo models.UserInfo
		if err := global.Db.Model(&models.User{}).Where("user_id =?", follow.FollowerID).First(&userInfo).Error; err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
		followedUserInfo = append(followedUserInfo, userInfo)
	}
	ResponseSuccess(c, gin.H{
		"followList":   followerUserInfo, // 你关注的用户列表
		"followedList": followedUserInfo, // 关注你的用户列表
	})

}

func IsFollow(c *gin.Context) {
	var follow models.FollowFilter
	if err := c.ShouldBindJSON(&follow); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if IsFollowerLogic(follow.FollowerID, follow.FollowedID) {
		ResponseSuccess(c, gin.H{
			"isFollow": true,
		})
		return
	}
	ResponseSuccess(c, gin.H{
		"isFollow": false,
	})
}

func IsFollowerLogic(followerId uint64, followedId uint64) bool {
	var follow models.Follow
	if err := global.Db.Where("follower_id = ? AND followed_id = ?", followerId, followedId).First(&follow).Error; err != nil {
		return false
	}
	return true
}
