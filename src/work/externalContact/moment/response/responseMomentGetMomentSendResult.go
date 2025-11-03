package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentSendResult struct {
	response.ResponseWork

	CustomerList []*power.HashMap `json:"customer_list"`
	NextCursor   string           `json:"next_cursor"`
}
