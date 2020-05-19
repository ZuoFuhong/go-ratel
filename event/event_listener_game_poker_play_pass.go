package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGamePokerPlayPass(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice(dataMap["clientNickname"].(string) + " passed. It is now " + dataMap["nextClientNickname"].(string) + "'s turn.")

	turnClientId := int(dataMap["nextClientId"].(float64))
	if ctx.UserId == turnClientId {
		ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
	}
}
