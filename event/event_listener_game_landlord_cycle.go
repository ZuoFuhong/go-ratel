package event

import "go-ratel/command"

func ListenerGameLandlordCycle(ctx *Context, data string) {
	command.PrintNotice("No player takes the landlord, so redealing cards.")
}
