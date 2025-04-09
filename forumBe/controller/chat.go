package controller

import (
	"forum/config"
	"forum/server"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	if fileName == "" {
		return
	}
	zap.L().Info("fileName", zap.String("fileName", fileName))
	data, _ := ioutil.ReadFile(config.AppConf.StaticFilePath + fileName)
	c.Writer.Write(data)
}

// 上传头像等文件
func Savefile(c *gin.Context) {
	namePreffix := uuid.New().String()

	userUuid := c.PostForm("uuid")

	file, _ := c.FormFile("file")
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePreffix + suffix

	zap.L().Info("file", zap.Any("file name", config.AppConf.StaticFilePath+newFileName))
	zap.L().Info("userUuid", zap.Any("userUuid name", userUuid))
	c.SaveUploadedFile(file, config.AppConf.StaticFilePath+newFileName)
	ResponseSuccessWithMsg(c, 1000, "success")
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 建立websocket连接
func RunSocekt(c *gin.Context) {
	user := c.Query("uid")
	if user == "" {
		return
	}
	zap.L().Info("newUser", zap.String("newUser", user))
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	client := &server.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}
	go func() {
		server.MyServer.Register <- client
	}()
	go client.Read()
	go client.Write()
}
