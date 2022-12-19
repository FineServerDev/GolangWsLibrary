package economy

const (
	Eco_Common_Error_Response = "common_error_response"
)

type CommonErrorResponse struct {
	//错误信息
	Message string `json:"message"`
}
