package models

import (
	"time"
)

type Message struct {
	ID          int32     `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   uint      `json:"deletedAt"`
	FromUserId  uint64    `json:"fromUserId" gorm:"index"`
	ToUserId    uint64    `json:"toUserId" gorm:"index;comment:'发送给端的id，可为用户id或者群id'"`
	Content     string    `json:"content" gorm:"type:varchar(2500)"`
	MessageType int16     `json:"messageType" gorm:"comment:'消息类型：1单聊，2群聊'"`
	ContentType int16     `json:"contentType" gorm:"comment:'消息内容类型：1文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天'"`
	Pic         string    `json:"pic" gorm:"type:text;comment:'缩略图"`
	Url         string    `json:"url" gorm:"type:varchar(350);comment:'文件或者图片地址'"`
}

type GroupMember struct {
	ID        int32     `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt uint      `json:"deletedAt"`
	UserId    int32     `json:"userId" gorm:"index;comment:'用户ID'"`
	GroupId   int32     `json:"groupId" gorm:"index;comment:'群组ID'"`
	Nickname  string    `json:"nickname" gorm:"type:varchar(350);comment:'昵称"`
	Mute      int16     `json:"mute" gorm:"comment:'是否禁言'"`
}

type Group struct {
	ID        int32     `json:"id" gorm:"primarykey"`
	Uuid      string    `json:"uuid" gorm:"type:varchar(150);not null;unique_index:idx_uuid;comment:'uuid'"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt uint16    `json:"deletedAt"`
	UserId    int32     `json:"userId" gorm:"index;comment:'群主ID'"`
	Name      string    `json:"name" gorm:"type:varchar(150);comment:'群名称"`
	Notice    string    `json:"notice" gorm:"type:varchar(350);comment:'群公告"`
}

type MessageRequest struct {
	MessageType    int32  `form:"messageType"`
	Uid            uint64 `form:"uid"`
	FriendUsername string `form:"friendUsername"`
	FUid           uint64 `form:"fUid"`
}

type MessageResponse struct {
	ID          int32     `json:"id" gorm:"primarykey"`
	FromUserId  uint64    `json:"fromUserId" gorm:"index"`
	ToUserId    uint64    `json:"toUserId" gorm:"index"`
	Content     string    `json:"content" gorm:"type:varchar(2500)"`
	ContentType int16     `json:"contentType" gorm:"comment:'消息内容类型：1文字，2语音，3视频'"`
	CreatedAt   time.Time `json:"createAt"`
	Url         string    `json:"url"`
}
