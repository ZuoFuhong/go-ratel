package event

import (
	"encoding/json"
	"go-ratel/command"
	"strings"
)

func ListenerGameLandlordElect(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)
	turnClientId := int(dataMap["nextClientId"].(float64))

	if dataMap["preClientNickname"] != nil {
		command.PrintNotice(dataMap["preClientNickname"].(string) + " don't rob the landlord!")
	}
	if turnClientId == ctx.UserId {
		command.PrintNotice("It's your turn. Do you want to rob the landlord? [Y/N] (enter [EXIT] to exit current room)")
		line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("Y/N")))
		switch line {
		case "EXIT":
			ctx.pushToServer(SERVER_CODE_CLIENT_EXIT, "")
		case "Y":
			ctx.pushToServer(CODE_GAME_LANDLORD_ELECT, "TRUE")
		case "N":
			ctx.pushToServer(CODE_GAME_LANDLORD_ELECT, "FALSE")
		default:
			command.PrintNotice("Invalid options")
			ListenerGameLandlordElect(ctx, data)
		}
	} else {
		command.PrintNotice("It's " + dataMap["nextClientNickname"].(string) + "'s turn. Please wait patiently for his/her confirmation !")
	}
}
