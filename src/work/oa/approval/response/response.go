package response

import "github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"

type ResponseCreateTemplate struct {
	response.ResponseWork

	TemplateId string `json:"template_id"`
}
