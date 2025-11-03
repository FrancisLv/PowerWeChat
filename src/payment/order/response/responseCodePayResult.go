package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseCodePayResult struct {
	response.ResponsePayment

	ResponseOrder
}
