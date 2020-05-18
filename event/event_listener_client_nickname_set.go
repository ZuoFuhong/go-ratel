package event

import (
	"fmt"
	"strconv"
)

const NICKNAME_MAX_LENGTH = 18

func ListenerClientNicknameSet(ctx *Context, data string) {
	fmt.Println("Please set your nickname (upto", strconv.Itoa(NICKNAME_MAX_LENGTH), "characters)")
	// todo:
}
