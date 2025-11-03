package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
