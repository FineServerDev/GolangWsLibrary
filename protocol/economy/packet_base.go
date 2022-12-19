package economy

type PacketBase struct {
	MessageType string `json:"message_type"`
	//PacketId    int         `json:"packet_id"`
	Data interface{} `json:"data"`
}
