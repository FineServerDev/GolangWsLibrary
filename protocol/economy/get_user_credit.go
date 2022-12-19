package economy

const (
	Eco_Get_Credit_Request  = "eco_get_user_credit_request"
	Eco_Get_Credit_Response = "eco_get_user_credit_response"
)

type GetUserCreditRequest struct {
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
}

type GetUserCreditResponse struct {
	//当前用户的金钱数量
	Credit int `json:"credit"`
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
}
