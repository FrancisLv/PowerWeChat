package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	response.ResponseWork

	States []*power.HashMap `json:"states"`
}
