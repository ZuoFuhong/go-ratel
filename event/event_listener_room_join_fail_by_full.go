package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomJoinFailByFull(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("Join room failed. Room " + strconv.Itoa(dataMap["roomId"].(int)) + " player count is full!")
	ListenerShowOptions(ctx, data)
}
