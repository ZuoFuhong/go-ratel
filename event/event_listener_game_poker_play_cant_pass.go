package event

import "go-ratel/command"

func ListenerGamePokerPlayCantPass(ctx *Context, data string) {
	command.PrintNotice("You played the previous card, so you can't pass.")
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
