package economy

const (
	Eco_Set_Credit_Request  = "eco_set_user_credit_request"
	Eco_Set_Credit_Response = "eco_set_user_credit_response"
)

type SetUserCredit struct {
	MessageType string `json:"message_type"`
	Data        struct {
		UserID string `json:"user_id"`
		Credit int    `json:"credit"`
	} `json:"data"`
}
