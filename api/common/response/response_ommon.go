package response

//Response is used for static shape json return
type SuccsessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildSuccsessResponse(message string, data interface{}) SuccsessResponse {
	res := SuccsessResponse{
		Message: message,
		Status:  "Success",
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(status string, message string) ErrorResponse {
	res := ErrorResponse{
		Status:  status,
		Message: message,
	}
	return res
}
