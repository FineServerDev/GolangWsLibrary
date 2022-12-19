package client

import (
	"GolangWsLibrary/protocol/economy"
	"GolangWsLibrary/utils"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	conn *websocket.Conn
}

func sendRequest[T interface{}](c *Client, jsonStruct interface{}) T {
	var result T
	c.conn.WriteJSON(jsonStruct)
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return result
	}
	var packetBase economy.PacketBase
	err = json.Unmarshal(message, &packetBase)
	if err != nil {
		return result
	}
	utils.MapToStruct(packetBase.Data, &result)
	return result
}

func StartConnect() {
	c, _, err := websocket.DefaultDialer.Dial("", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	client := &Client{conn: c}

	a1 := sendRequest[economy.SetUserCreditResponse](client, economy.PacketBase{
		MessageType: economy.Eco_Set_Credit_Request,
		Data: economy.SetUserCreditRequest{
			UserID: "1",
			Credit: 100,
		},
	})

	log.Println("id", a1.UserID, "Credit", a1.Credit)

	for {
	}
}
