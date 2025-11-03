package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGAICrop struct {
	response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}
