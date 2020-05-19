package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameStarting(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("Game starting !!")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &pokers)

	command.PrintNotice("")
	command.PrintNotice("Your pokers are")
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	ListenerGameLandlordElect(ctx, data)
}
