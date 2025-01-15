package models

type User struct {
	UserId          uint64 `json:"user_id"`                               // 用户Id
	UserName        string `json:"username" binding:"required"`           // 用户名
	Email           string `json:"email" binding:"required"`              // 邮箱
	Gender          int    `json:"gender" binding:"oneof=0 1 2"`          // 性别 0:未知 1:男 2:女
	Password        string `json:"password,omitempty" binding:"required"` // 密码
	ConfirmPassword string `json:"confirm_password,omitempty" binding:"required,eqfield=Password"`
}
