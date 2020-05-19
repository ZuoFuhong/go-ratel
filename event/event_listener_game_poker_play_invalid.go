package event

import "go-ratel/command"

func ListenerGamePokerPlayInvalid(ctx *Context, data string) {
	command.PrintNotice("Out pokers' format is invalid")
	if ctx.LastPokers != nil {
		command.PrintNotice(ctx.LastSellClientNickname + "[" + ctx.LastSellClientType + "] played:")
		command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
	}
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
