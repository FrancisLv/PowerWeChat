package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyList struct {
	response.ResponseWork

	StrategyID int64 `json:"strategy_id"`
}
