package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	response.ResponseWork

	LivingInfo *power.HashMap `json:"living_info"`
}
