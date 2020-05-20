package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomJoinSuccess(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	ctx.InitLastSellInfo()
	joinClientId := int(dataMap["clientId"].(float64))

	if ctx.UserId == joinClientId {
		command.PrintNotice("You have joined room：" + strconv.Itoa(int(dataMap["roomId"].(float64))) + ". There are " + strconv.Itoa(int(dataMap["roomClientCount"].(float64))) + " players in the room now.")
		command.PrintNotice("Please wait for other players to join, start a good game when the room player reaches three !")
	} else {
		command.PrintNotice(dataMap["clientNickname"].(string) + " joined room, the current number of room player is " + strconv.Itoa(int(dataMap["roomClientCount"].(float64))))
	}
}
