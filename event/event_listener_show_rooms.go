package event

import (
	"encoding/json"
	"fmt"
	"go-ratel/command"
	"strconv"
)

func ListenerShowRooms(ctx *Context, data string) {
	roomList := make([]map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(data), &roomList)

	if len(roomList) > 0 {
		format := "#\t%s\t|\t%-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s\t|\t%-6s\t|\t%-6s\t#"
		command.PrintNotice(fmt.Sprintf(format, "ID", "OWNER", "COUNT", "TYPE"))
		for _, room := range roomList {
			command.PrintNotice(fmt.Sprintf(format, strconv.Itoa(int(room["roomId"].(float64))), room["roomOwner"].(string), strconv.Itoa(int(room["roomClientCount"].(float64))), room["roomType"].(string)))
		}
		command.PrintNotice("")
		ListenerShowOptionsPVP(ctx, data)
	} else {
		command.PrintNotice("No available room, please create a room ！")
		ListenerShowOptionsPVP(ctx, data)
	}
}
