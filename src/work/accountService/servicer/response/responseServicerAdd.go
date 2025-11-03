package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseServicerAdd struct {
	response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
