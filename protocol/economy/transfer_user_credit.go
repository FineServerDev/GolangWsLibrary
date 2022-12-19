package economy

const (
	Eco_Transfer_Credit_Request  = "eco_transfer_user_credit_request"
	Eco_Transfer_Credit_Response = "eco_transfer_user_credit_response"
)

type TransferUserCreditRequest struct {
	//转出用户唯一标识符（可能是UUID/XUID)
	FromUserID string `json:"from_user_id"`
	//转入用户唯一标识符（可能是UUID/XUID)
	ToUserID string `json:"to_user_id"`
	//转账的金钱数量
	Credit int `json:"credit"`
}

type TransferUserCreditResponse struct {
	//转出用户的金钱数量
	FromUserCredit int `json:"from_user_credit"`
	//转出用户唯一标识符（可能是UUID/XUID)
	FromUserID string `json:"from_user_id"`
	//转入用户的金钱数量
	ToUserCredit int `json:"to_user_credit"`
	//转入用户唯一标识符（可能是UUID/XUID)
	ToUserID string `json:"to_user_id"`
}
