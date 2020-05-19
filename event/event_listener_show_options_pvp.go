package event

import (
	"go-ratel/command"
	"strconv"
	"strings"
)

func ListenerShowOptionsPVP(ctx *Context, data string) {
	command.PrintNotice("PVP: ")
	command.PrintNotice("1. Create Room")
	command.PrintNotice("2. Room List")
	command.PrintNotice("3. Join Room")
	command.PrintNotice("Please enter the number of options (enter [BACK] return options list)")
	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("pvp")))
	if line == "BACK" {
		ListenerShowOptions(ctx, data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ctx.pushToServer(SERVER_CODE_ROOM_CREATE, "")
		case 2:
			ctx.pushToServer(SERVER_CODE_GET_ROOMS, "")
		case 3:
			command.PrintNotice("Please enter the room id you want to join (enter [BACK] return options list)")
			line := command.DeletePreAndSufSpace(command.Write("roomid"))
			if strings.ToUpper(line) == "BACK" {
				ListenerShowOptionsPVP(ctx, data)
			} else {
				roomid, e := strconv.Atoi(line)
				if e != nil {
					roomid = -1
				}
				if roomid < 1 {
					command.PrintNotice("Invalid options, please choose again：")
					ListenerShowOptionsPVP(ctx, data)
				} else {
					ctx.pushToServer(SERVER_CODE_ROOM_JOIN, strconv.Itoa(roomid))
				}
			}
		default:
			command.PrintNotice("Invalid option, please choose again：")
			ListenerShowOptionsPVP(ctx, data)
		}
	}
}
