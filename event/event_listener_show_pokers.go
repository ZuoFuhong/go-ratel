package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerShowPokers(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	ctx.LastSellClientNickname = dataMap["clientNickname"].(string)
	ctx.LastSellClientType = dataMap["clientType"].(string)

	command.PrintNotice(ctx.LastSellClientNickname + "[" + ctx.LastSellClientType + "] played:")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal(pokersBytes, &pokers)
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	if dataMap["sellClinetNickname"] != nil {
		command.PrintNotice("Next player is " + dataMap["sellClinetNickname"].(string) + ". Please wait for him to play his pokers.")
	}
}
