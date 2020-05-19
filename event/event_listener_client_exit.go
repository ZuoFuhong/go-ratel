package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerClientExit(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	exitClientId := int(dataMap["exitClientId"].(float64))
	var role string
	if exitClientId == ctx.UserId {
		role = "You"
	} else {
		role = dataMap["exitClientNickname"].(string)
	}
	command.PrintNotice(role + " exit from the room. Room disbanded!!\n")
	ListenerShowOptions(ctx, data)
}
