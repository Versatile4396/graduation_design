package server

import (
	"encoding/base64"
	"forum/config"
	"forum/constant"
	"forum/dao/service"
	util "forum/pkg/common"
	"forum/pkg/protocol"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

var MyServer = NewServer()

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {
	zap.L().Info("start server", zap.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			zap.L().Info("login", zap.Any("login", "new user login in"+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			zap.L().Info("loginout", zap.Any("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// 一般消息，比如文本消息，视频文件消息等
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
					// 保存消息只会在存在socket的一个端上进行保存，防止分布式部署后，消息重复问题
					_, exits := s.Clients[msg.From]
					if exits {
						saveMessage(msg)
					}

					if msg.MessageType == constant.MESSAGE_TYPE_USER {
						client, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err == nil {
								client.Send <- msgByte
							}
						}
					} else if msg.MessageType == constant.MESSAGE_TYPE_GROUP {
						// sendGroupMessage(msg, s)
					}
				} else {
					// 语音电话，视频电话等，仅支持单人聊天，不支持群聊
					// 不保存文件，直接进行转发
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}

			} else {
				// 无对应接受人员进行广播
				for id, conn := range s.Clients {
					zap.L().Info("allUser", zap.Any("allUser", id))

					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}

// 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(message *protocol.Message) {
	// 如果上传的是base64字符串文件，解析文件保存
	if message.ContentType == 2 {
		url := uuid.New().String() + ".png"
		index := strings.Index(message.Content, "base64")
		index += 7

		content := message.Content
		content = content[index:]

		dataBuffer, dataErr := base64.StdEncoding.DecodeString(content)
		if dataErr != nil {
			zap.L().Error("transfer base64 to file error", zap.String("transfer base64 to file error", dataErr.Error()))
			return
		}
		err := ioutil.WriteFile(config.AppConf.StaticFilePath+url, dataBuffer, 0666)
		if err != nil {
			zap.L().Error("write file error", zap.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.Content = ""
	} else if message.ContentType == 3 {
		// 普通的文件二进制上传
		fileSuffix := util.GetFileType(message.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(message.FileSuffix)
		}
		contentType := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := ioutil.WriteFile(config.AppConf.StaticFilePath+url, message.File, 0666)
		if err != nil {
			zap.L().Error("write file error", zap.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}

	service.MessageService.SaveMessage(*message)
}
