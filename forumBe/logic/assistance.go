package logic

import (
	"fmt"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

func CreateAssistance(a *models.Assistance) (ra *models.Assistance, err error) {
	a.AssistanceId, err = snowflake.GetID()
	if err != nil {
		return nil, err
	}
	err = global.Db.Create(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func DeleteAssistance(postId uint64) (err error) {
	err = global.Db.Delete(&models.Assistance{}, postId).Error
	return
}

func UpdateAssistance(a *models.Assistance) (err error) {
	err = global.Db.Model(&models.Assistance{}).Where("assistance_id = ?", a.AssistanceId).Updates(a).Error
	return err
}

func GetAssistanceList(f models.AssistanceFilter) (rAssistance []*models.Assistance, err error) {
	if f.Pagination == nil {
		f.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	offset := (f.Pagination.Page - 1) * f.Pagination.PageSize
	if f.Title == "" {
		f.Title = "%%"
	}
	query := global.Db.Model(&rAssistance)
	if f.Title != "" {
		query = query.Where("title LIKE?", "%"+f.Title+"%")
	}
	if f.CategoryId != 0 {
		query = query.Where("category_id =?", f.CategoryId)
	}

	if f.UserId != 0 {
		query = query.Where("user_id =?", f.UserId)
	}

	orderDir := "desc"
	if f.OrderByTime {
		orderDir = "asc"
	}
	query = query.Order(fmt.Sprintf("created_at %s", orderDir)).Offset(offset).Limit(f.Pagination.PageSize)
	query.Find(&rAssistance)
	var userInfo models.UserInfo
	for _, v := range rAssistance {
		userInfo, err = GetUserInfo(v.UserId)
		if err != nil {
			return nil, err
		}
		v.UserInfo = &userInfo
	}
	return rAssistance, nil
}
func GetUserInfo(userId uint64) (userInfo models.UserInfo, err error) {
	err = global.Db.Model(&models.User{}).Where("user_id =?", userId).First(&userInfo).Error
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func CreateAssistanceComment(f *models.AssistanceComment) (rf *models.AssistanceComment, err error) {
	f.CommentId, err = snowflake.GetID()
	if err != nil {
		return nil, err
	}

	err = global.Db.Create(&f).Error
	if err != nil {
		return nil, err
	}
	return rf, nil
}

func DeleteAssistanceComment(postId uint64) (err error) {
	err = global.Db.Delete(&models.AssistanceComment{}, postId).Error
	return
}

func GetAssistanceCommentList(f models.AssistanceCommentFilter) (rAssistanceComment []*models.AssistanceComment, err error) {
	if f.Pagination == nil {
		f.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	offset := (f.Pagination.Page - 1) * f.Pagination.PageSize
	query := global.Db.Model(&rAssistanceComment)
	if f.AssistanceId != 0 {
		query = query.Where("assistance_id =?", f.AssistanceId)
	}
	if f.UserId != 0 {
		query = query.Where("user_id =?", f.UserId)
	}
	query = query.Offset(offset).Limit(f.Pagination.PageSize)
	query.Find(&rAssistanceComment)
	for _, v := range rAssistanceComment {
		var userInfo models.UserInfo
		userInfo, err = GetUserInfo(uint64(v.UserId))
		v.UserInfo = &userInfo
	}
	return
}
