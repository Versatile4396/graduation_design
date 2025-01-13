package controller

import "github.com/gin-gonic/gin"

// Ping godoc
// @Summary      用户注册
// @Description  用户注册接口
// @Success      200  {string}  string
// @body        name body     string  true  "name"
// @Produce      json
// @Router       /register [post]
func UserRegisterController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}

func UserLoginController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}

func UserUpdateController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
func UserDeleteController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
func UserInfoController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
