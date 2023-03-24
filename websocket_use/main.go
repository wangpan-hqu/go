package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)

/*
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}
func checkOrigin(r *http.Request) bool {
	return true
}
var conn *websocket.Conn
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//判断请求是否为websocket升级请求。
	if websocket.IsWebSocketUpgrade(r) {
		conn, _ := upgrader.Upgrade(w, r, w.Header())
		conn.WriteMessage(websocket.TextMessage, []byte("wxm.alming"))
		go func() {
			for {
				t, c, _ := conn.ReadMessage()
				fmt.Println(t, string(c))
				if t == -1 {
					return
				}
			}
		}()
	} else {

	}
}
*/
/*
func (c *websocket.Conn) SetCloseHandler(h func(code int, text string) error) {
	if h == nil {
		h = func(code int, text string) error {
			message := FormatCloseMessage(code, "")
			c.WriteControl(CloseMessage, message, time.Now().Add(writeWait))
			return nil
		}
	}
	c.handleClose = h
}
*/

type msgJson struct {
	M_str string
}

func main() {
	//创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:18802")

	//控制连接我的客户端数量
	var num = 0

	//打印错误
	fmt.Println(err)

	mux := http.NewServeMux()
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//升级成功后将获取到WebSocket.Conn 利用这个Conn可进行消息收发
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("upgrade error")
			return
		}

		num++

		//将msg转化成字节数组
		msg := msgJson{M_str: "xxxxx"}
		msgbyte, err := json.Marshal(msg)

		//var kkk msgJson
		//json.Unmarshal(msgbyte,&kkk)

		if err == nil {
			//发送
			c.WriteMessage(websocket.BinaryMessage, msgbyte)
		}
	})

	sv := http.Server{Addr: "127.0.0.1:18802", Handler: mux}

	go func() {
		err := sv.Serve(listener)
		fmt.Println("Http.serve error:", err)
	}()

	//阻塞主线程
	for {
		if num > 100 {
			time.Sleep(1000)
			break
		}
	}
}
