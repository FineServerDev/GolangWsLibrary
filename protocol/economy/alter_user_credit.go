package economy

const (
	Eco_Alter_Credit_Request  = "eco_alter_user_credit_request"
	Eco_Alter_Credit_Response = "eco_alter_user_credit_response"
)

type AlterUserCreditRequest struct {
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
	//修改的金钱数量
	Credit int `json:"credit"`
}

type AlterUserCreditResponse struct {
	//当前用户的金钱数量
	Credit int `json:"credit"`
	//用户唯一标识符（可能是UUID/XUID)
	UserID string `json:"user_id"`
}

func CreateAlterUserCreditRequest(UserID string, Credit int) *PacketBase {
	return &PacketBase{
		MessageType: Eco_Alter_Credit_Request,
		Data: &AlterUserCreditRequest{
			UserID: UserID,
			Credit: Credit,
		},
	}
}
