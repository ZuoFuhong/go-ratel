package event

import (
	"go-ratel/command"
	"strconv"
)

func ListenerClientConnect(ctx *Context, data string) {
	command.PrintNotice("Connection to server is successful. Welcome to ratel!!")
	ctx.UserId, _ = strconv.Atoi(data)
}
