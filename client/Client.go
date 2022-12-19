package client

import "github.com/gorilla/websocket"

func StartConnect() {
	//连接wss服务器 wss:/3.ifine.eu:10000/socket

	dial, _, err := websocket.DefaultDialer.Dial("wss://3.ifine.eu:10000/socket", nil)
	if err != nil {
		return
	}

}
