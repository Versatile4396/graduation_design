package server

import (
	"fmt"
	"forum/config"
	"forum/constant"
	"forum/pkg/kafka"
	"forum/pkg/protocol"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			zap.L().Error("client read message error", zap.Any("client read message error", err.Error()))
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}
		msg := &protocol.Message{}
		fmt.Println("proto", message)
		err = proto.Unmarshal(message, msg)
		if err != nil {
			zap.L().Error("client unmarshal message error", zap.Any("client unmarshal message error", err.Error()))
		}
		fmt.Println("message", message)
		fmt.Println("msg", msg, msg.Content, msg.Type)
		// pong
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err2 := proto.Marshal(pong)
			if nil != err2 {
				zap.L().Error("client marshal message error", zap.Any("client marshal message error", err2.Error()))
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			if config.AppConf.MsgChannelType.ChannelType == constant.KAFKA {
				kafka.Send(message)
			} else {
				MyServer.Broadcast <- message
			}
		}
	}
}
func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
