package client

import (
	"GolangWsLibrary/protocol/economy"
	"GolangWsLibrary/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

const URL = "wss://sh3.ifine.eu:10000/socket"

var GlobalClient *Client

type Client struct {
	conn *websocket.Conn
}

func SendRequest[T interface{}](jsonStruct interface{}) (T, error) {
	var result T
	GlobalClient.conn.WriteJSON(jsonStruct)
	_, message, err := GlobalClient.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return result, err
	}
	var packetBase economy.PacketBase
	err = json.Unmarshal(message, &packetBase)
	if err != nil {
		return result, err
	}
	if packetBase.MessageType == "common_error_response" {
		var commonErrorResponse economy.CommonErrorResponse
		err = utils.MapToStruct(packetBase.Data, &commonErrorResponse)
		return result, fmt.Errorf(commonErrorResponse.Message)
	}
	utils.MapToStruct(packetBase.Data, &result)
	fmt.Println(string(message))
	return result, nil
}

func StartConnect() {
	c, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	client := &Client{conn: c}
	GlobalClient = client
}
