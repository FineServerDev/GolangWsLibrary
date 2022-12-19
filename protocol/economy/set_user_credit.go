package economy

const (
	Eco_Set_Credit_Request  = "eco_set_user_credit_request"
	Eco_Set_Credit_Response = "eco_set_user_credit_response"
)

type SetUserCreditRequest struct {
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
	//设置的金钱数量
	Credit int `json:"credit"`
}

type SetUserCreditResponse struct {
	//当前用户的金钱数量
	Credit int `json:"credit"`
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
}

func CreateSetUserCreditRequest(UserID string, Credit int) *PacketBase {
	return &PacketBase{
		MessageType: Eco_Set_Credit_Request,
		Data: &SetUserCreditRequest{
			UserID: UserID,
			Credit: Credit,
		},
	}
}
