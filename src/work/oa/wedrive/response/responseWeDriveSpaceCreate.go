package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}
