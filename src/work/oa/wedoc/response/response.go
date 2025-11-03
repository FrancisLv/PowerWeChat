package response

import "github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"

type ResponseWeDocCreateForm struct {
	response.ResponseWork
	FormId string `json:"formid"`
}
