package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerRoomJoinFailByInExist(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("Join room failed. Room " + dataMap["roomId"].(string) + " inexists!")
	ListenerShowOptions(ctx, data)
}
