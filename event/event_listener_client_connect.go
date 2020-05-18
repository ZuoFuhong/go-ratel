package event

import (
	"fmt"
	"strconv"
)

func ListenerClientConnect(ctx *Context, data string) {
	fmt.Println("Connection to server is successful. Welcome to ratel!!")
	ctx.UserId, _ = strconv.Atoi(data)
}
