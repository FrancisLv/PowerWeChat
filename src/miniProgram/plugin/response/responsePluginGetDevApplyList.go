package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
