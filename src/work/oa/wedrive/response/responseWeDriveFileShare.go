package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileShare struct {
	response.ResponseWork

	ShareURL string `json:"share_url"`
}
