package controller

import (
	"forum/config"
	"forum/dao/service"
	"forum/logic"
	"forum/models"
	"forum/server"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func GetFile(c *gin.Context) {
	fileName := c.Param("filename")

	zap.L().Info("filename", zap.String("filename", fileName))
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

func GetMessage(c *gin.Context) {
	zap.L().Info(c.Query("uid"))
	var messageRequest models.MessageRequest
	err := c.BindQuery(&messageRequest)
	if nil != err {
		zap.L().Error("bindQueryError", zap.Any("bindQueryError", err))
	}
	zap.L().Info("messageRequest params: ", zap.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "获取历史消息失败")
		return
	}
	// 拿用户信息
	toUserInfo, err := logic.GetUserInfo(messageRequest.ToUid)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "获取用户信息内容失败")
		return
	}
	type TRes struct {
		MessageInfo []models.MessageResponse `json:"messageInfo"`
		ToUserInfo  models.UserInfo          `json:"toUserInfo"`
	}

	ResponseSuccess(c, &TRes{
		MessageInfo: messages,
		ToUserInfo:  toUserInfo,
	})
}

func GetFriend(c *gin.Context) {
	uid := c.Query("uid")
	postId, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		return
	}
	if uid == "" {
		return
	}
	friends, err := logic.GetFriendList(postId)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "获取聊天列表内容失败")
		return
	}
	ResponseSuccess(c, friends)
}
