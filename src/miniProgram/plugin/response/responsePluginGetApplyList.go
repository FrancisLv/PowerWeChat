package response

import (
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/power"
	"github.com/FrancisLv/PowerWeChat/v3/src/kernel/response"
)

type ResponsePluginGetPluginList struct {
	response.ResponseMiniProgram
	PluginList []*power.HashMap `json:"plugin_list"`
}
