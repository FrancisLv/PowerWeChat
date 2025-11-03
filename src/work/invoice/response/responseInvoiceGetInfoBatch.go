package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseInvoiceGetInfoBatch struct {
	response.ResponseWork

	ItemList []*power.HashMap `json:"item_list"`
}
