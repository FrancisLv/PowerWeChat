package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	response.ResponseMiniProgram

	UserName string `json:"username"`
}
