package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}
