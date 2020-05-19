package event

import "go-ratel/command"

func ListenerGamePokerPlayLess(ctx *Context, data string) {
	command.PrintNotice("Your pokers' type has lower rank than the previous. You could not play this combination !!")
	if ctx.LastPokers != nil {
		command.PrintNotice(ctx.LastSellClientNickname + "[" + ctx.LastSellClientType + "] played:")
		command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
	}
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
