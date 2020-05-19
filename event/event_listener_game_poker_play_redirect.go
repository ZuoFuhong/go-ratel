package event

import (
	"encoding/json"
	"fmt"
	"go-ratel/command"
	"strconv"
	"strings"
)

var choose = [...]string{"UP", "DOWN"}
var format = "\n[%-4s] %-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s  surplus %-2s [%-8s]"

func ListenerGamePokerPlayRedirect(ctx *Context, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	sellClientId := int(dataMap["sellClientId"].(float64))

	clientInfos := make([]map[string]interface{}, 0)
	clientInfoBytes, _ := json.Marshal(dataMap["clientInfos"])
	_ = json.Unmarshal(clientInfoBytes, &clientInfos)

	for index := 0; index < 2; index++ {
		for _, clientInfo := range clientInfos {
			position := clientInfo["position"].(string)
			if strings.ToUpper(position) == strings.ToUpper(choose[index]) {
				command.PrintNotice(fmt.Sprintf(format, clientInfo["position"].(string), clientInfo["clientNickname"].(string), strconv.Itoa(int(clientInfo["surplus"].(float64))), clientInfo["type"].(string)))
			}
		}
	}
	command.PrintNotice("")
	if sellClientId == ctx.UserId {
		ListenerGamePokerPlay(ctx, data)
	} else {
		command.PrintNotice("Next player is " + dataMap["sellClinetNickname"].(string) + ". Please wait for him to play his cards.")
	}
}
