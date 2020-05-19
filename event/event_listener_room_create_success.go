package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerRoomCreateSuccess(ctx *Context, data string) {
	room := Room{}
	_ = json.Unmarshal([]byte(data), &room)

	ctx.InitLastSellInfo()

	command.PrintNotice("You have created a room with id " + strconv.Itoa(room.Id))
	command.PrintNotice("Please wait for other players to join !")
}
