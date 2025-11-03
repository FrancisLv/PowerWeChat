package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseMenuCreate struct {
	response.ResponseOfficialAccount
}

type ResponseMenuCreateConditional struct {
	response.ResponseOfficialAccount

	MenuID string `json:"menuid"`
}
