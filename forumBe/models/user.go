package models

type User struct {
	UserId          uint64 `json:"user_id"`                               // 用户Id
	UserName        string `json:"username" binding:"required"`           // 用户名
	Email           string `json:"email" binding:"required"`              // 邮箱
	Gender          int    `json:"gender" binding:"oneof=0 1 2 3"`        // 性别 0:未知 1:男 2:女
	Password        string `json:"password,omitempty" binding:"required"` // 密码
	ConfirmPassword string `json:"confirm_password,omitempty" binding:"required,eqfield=Password"`
	Avatar          string `json:"avatar"`
	AccessToken     string `json:"access_token" gorm:"-"`
	RefreshToken    string `json:"refresh_token" gorm:"-"`
}
type UserInfo struct {
	UserId   uint64 `json:"user_id"`                        // 用户Id
	UserName string `json:"username" binding:"required"`    // 用户名
	Email    string `json:"email" binding:"required"`       // 邮箱
	Gender   int    `json:"gender" binding:"oneof=0 1 2 3"` // 性别 0:未知 1:男 2:女
	Avatar   string `json:"avatar"`
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCountInfo struct {
	ViewCount    int64 `json:"view_count"`
	LikeCount    int64 `json:"like_count"`
	AritcleCount int64 `json:"article_count"`
	CommentCount int64 `json:"comment_cnt"`
}
