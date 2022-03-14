package helpers

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//func SuccessfullyGet(data interface{}) interface{} {
//	response := make(map[string]interface{})
//	response["success"] = true
//	response["message"] = "successfully return"
//	response["data"] = data
//	return response
//}
