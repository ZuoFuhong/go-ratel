package event

import "go-ratel/command"

func ListenerPVEDifficultyNotSupport(ctx *Context, data string) {
	command.PrintNotice("The current difficulty coefficient is not supported, please pay attention to the following.\n")
	ListenerShowOptionsPVE(ctx, data)
}
