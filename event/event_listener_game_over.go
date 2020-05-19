package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameOver(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("\nPlayer " + dataMap["winnerNickname"].(string) + "[" + dataMap["winnerType"].(string) + "]" + " won the game")
	command.PrintNotice("Game over, friendship first, competition second\n")
}
