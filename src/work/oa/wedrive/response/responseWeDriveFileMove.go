package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileMove struct {
	response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
