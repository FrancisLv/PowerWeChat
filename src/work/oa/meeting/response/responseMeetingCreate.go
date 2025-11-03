package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingCreate struct {
	response.ResponseWork

	MeetingID int `json:"meetingid"`
}
