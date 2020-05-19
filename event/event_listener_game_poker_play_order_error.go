package event

import "go-ratel/command"

func ListenerGamePokerPlayOrderError(ctx *Context, data string) {
	command.PrintNotice("Not turn you to operate, please wait other player !!")
}
