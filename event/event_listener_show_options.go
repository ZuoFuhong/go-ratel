package event

import (
	"go-ratel/command"
	"os"
	"strconv"
	"strings"
)

func ListenerShowOptions(ctx *Context, data string) {
	command.PrintNotice("Options: ")
	command.PrintNotice("1. PvP")
	command.PrintNotice("2. PvE")
	command.PrintNotice("3. Setting")
	command.PrintNotice("Please enter the number of options (enter [EXIT] log out)")

	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("options")))
	if line == "EXIT" {
		os.Exit(0)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ListenerShowOptionsPVP(ctx, data)
		case 2:
			ListenerShowOptionsPVE(ctx, data)
		case 3:
			ListenerShowOptionsSettings(ctx, data)
		default:
			command.PrintNotice("Invalid option, please choose again：")
			ListenerShowOptions(ctx, data)
		}
	}
}
