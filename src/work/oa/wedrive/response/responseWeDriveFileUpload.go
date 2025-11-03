package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileUpload struct {
	response.ResponseWork

	FileID string `json:"fileid"`
}
