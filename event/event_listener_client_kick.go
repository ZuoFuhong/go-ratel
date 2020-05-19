package event

import "go-ratel/command"

func ListenerClientKick(ctx *Context, data string) {
	command.PrintNotice("As a result of long time do not operate, be forced by the system to kick out of the room\n")
	ListenerShowOptions(ctx, data)
}
