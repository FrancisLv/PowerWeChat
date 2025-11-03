package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
