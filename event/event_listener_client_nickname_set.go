package event

import (
	"encoding/json"
	"go-ratel/command"
	"strconv"
)

func ListenerClientNicknameSet(ctx *Context, data string) {
	if data == "" {
		dataMap := make(map[string]interface{})
		if dataMap["invalidLength"] != nil {
			command.PrintNotice("Your nickname length was invalid: " + strconv.Itoa(dataMap["invalidLength"].(int)))
		}
	}

	command.PrintNotice("Please set your nickname (upto " + strconv.Itoa(NICKNAME_MAX_LENGTH) + " characters)")
	nickname := command.DeletePreAndSufSpace(command.Write("nickname"))
	if len(nickname) > NICKNAME_MAX_LENGTH {
		result := make(map[string]interface{})
		result["invalidLength"] = len(nickname)
		resultJson, _ := json.Marshal(&result)
		ListenerClientNicknameSet(ctx, string(resultJson))
	} else {
		ctx.pushToServer(SERVER_CODE_CLIENT_NICKNAME_SET, nickname)
	}
}
