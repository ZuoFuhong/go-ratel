package event

import (
	"encoding/json"
	"go-ratel/command"
)

func ListenerGameLandlordConfirm(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	landlordNickname := dataMap["landlordNickname"].(string)
	command.PrintNotice(landlordNickname + " grabbed the landlord and got extra three poker shots")

	// 序列化
	additionalPokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["additionalPokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &additionalPokers)

	command.PrintPokers(additionalPokers, ctx.PokerPrinterType)
	ctx.pushToServer(CODE_GAME_POKER_PLAY_REDIRECT, "")
}
