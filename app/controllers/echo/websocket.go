package echo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

//upGrader 设置websocket
var upGrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool { //CheckOrigin防止跨站点的请求伪造
		return true
	},
}

// The message types are defined in RFC 6455, section 11.8.
const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

type ResponseMsg struct {
	RemoteAddr  net.Addr    `json:"remote_addr"`
	MessageType int         `json:"message_type"`
	Message     interface{} `json:"message"`
	Err         error       `json:"err"`
}

func Websocket(c *gin.Context) {
	//升级get请求为websocket协议
	connWs, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("ws.upgrade Err: %v\n", err)
		return
	}

	defer func(connWs *websocket.Conn) {
		if err := connWs.Close(); err != nil {
			log.Printf("ws.close Err: %v\n", err)
		}
	}(connWs) //返回前关闭

	for {
		var msgData ResponseMsg
		//读取ws中的数据
		msgType, msg, err := connWs.ReadMessage()
		if err != nil {
			log.Printf("ws.ReadMessage to error: %v", err)
			break
		}

		if msgType == TextMessage {

			msgData.RemoteAddr = connWs.RemoteAddr()
			msgData.MessageType = msgType
			msgData.Err = err

			//验证消息是否是json
			if json.Valid(msg) {
				fmt.Println("ok json")
				if err = json.Unmarshal(msg, &msgData.Message); err != nil {
					log.Printf("ws.ReadMessage to json.Unmarshal error: %v", err)
				}
			} else {
				msgData.Message = string(msg)
			}

			message, _ := json.Marshal(msgData)

			if err = connWs.WriteMessage(msgType, message); err != nil {
				log.Printf("ws.WriteMessage to error: %v\n", err)
				break
			}
		}

		log.Println("ws.WriteMessage.ok")
	}
}
