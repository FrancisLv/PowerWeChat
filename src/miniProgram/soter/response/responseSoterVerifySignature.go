package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
