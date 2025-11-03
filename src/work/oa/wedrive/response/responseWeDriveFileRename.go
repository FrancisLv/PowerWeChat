package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
