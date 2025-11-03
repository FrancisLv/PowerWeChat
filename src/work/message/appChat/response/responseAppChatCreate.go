package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatCreate struct {
	response.ResponseWork

	ChatID string `json:"chatid"`
}
