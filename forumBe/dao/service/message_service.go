package service

import (
	"forum/constant"
	"forum/global"
	"forum/models"
	"forum/pkg/protocol"

	"go.uber.org/zap"
)

const NULL_ID uint64 = 0

type messageService struct {
}

var MessageService = new(messageService)

func (m *messageService) SaveMessage(message protocol.Message) {
	db := global.Db
	var fromUser models.User
	db.Find(&fromUser, "user_id = ?", message.From)
	if NULL_ID == fromUser.UserId {
		zap.L().Error("SaveMessage not find from user", zap.Any("SaveMessage not find from user", fromUser.UserId))
		return
	}

	var toUserId uint64 = 0

	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var toUser models.User
		db.Find(&toUser, "user_id = ?", message.To)
		if NULL_ID == toUser.UserId {
			return
		}
		toUserId = toUser.UserId
	}

	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		var group models.Group
		db.Find(&group, "user_id = ?", message.To)
		if NULL_ID == uint64(group.ID) {
			return
		}
		toUserId = uint64(group.ID)
	}

	saveMessage := models.Message{
		FromUserId:  fromUser.UserId,
		ToUserId:    toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	db.Save(&saveMessage)
}
