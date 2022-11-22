package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
type EmptyObj struct {
}

func BuildResponse(status bool, message string, errors interface{}, data interface{}) Response {
	var res Response
	res.Status = status
	res.Message = message
	res.Errors = errors
	res.Data = data
	return res
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
