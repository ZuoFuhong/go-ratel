package command

import (
	"fmt"
	"go-ratel/common"
)

type Poker = common.Poker
type PokerEntity = common.PokerEntity

func PrintPokers(pokers []Poker, printerType int) {
	pokerEntitys := common.ConvertEnumPoker(&pokers)
	sortPokers(pokerEntitys)
	switch printerType {
	case 0:
		buildHandStringSharp(*pokerEntitys)
	case 1:
		buildHandStringRounded(*pokerEntitys)
	case 2:
		textOnly(*pokerEntitys)
	case 3:
		textOnlyNoType(*pokerEntitys)
	default:
		buildHandStringSharp(*pokerEntitys)
	}
}

func sortPokers(pokers *[]PokerEntity) {
	innerPokers := *pokers
	length := len(innerPokers)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {

			if innerPokers[j].Level.Level > innerPokers[j+1].Level.Level {
				temp := innerPokers[j]
				innerPokers[j] = innerPokers[j+1]
				innerPokers[j+1] = temp
			}
		}
	}
	*pokers = innerPokers
}

func buildHandStringSharp(pokerEntitys []PokerEntity) {
	outputStr := ""
	if pokerEntitys != nil && len(pokerEntitys) > 0 {
		for i := 0; i < len(pokerEntitys); i++ {
			if i == 0 {
				outputStr += "┌──┐"
			} else {
				outputStr += "──┐"
			}
		}
	}
	outputStr += "\n"
	for i := 0; i < len(pokerEntitys); i++ {
		if i == 0 {
			outputStr += "|"
		}
		name := pokerEntitys[i].Level.Name
		if len(name) == 1 {
			outputStr += name + " " + "|"
		} else {
			outputStr += name + "|"
		}
	}
	outputStr += "\n"
	for i := 0; i < len(pokerEntitys); i++ {
		if i == 0 {
			outputStr += "|"
		}
		outputStr += pokerEntitys[i].Type + " |"
	}
	outputStr += "\n"
	for i := 0; i < len(pokerEntitys); i++ {
		if i == 0 {
			outputStr += "└──┘"
		} else {
			outputStr += "──┘"
		}
	}
	fmt.Println(outputStr)
}

func buildHandStringRounded(pokerEntitys []PokerEntity) {
	outputStr := ""
	if pokerEntitys != nil && len(pokerEntitys) > 0 {
		for i := 0; i < len(pokerEntitys); i++ {
			if i == 0 {
				outputStr += "┌──╮"
			} else {
				outputStr += "──╮"
			}
		}
		outputStr += "\n"
		for i := 0; i < len(pokerEntitys); i++ {
			if i == 0 {
				outputStr += "|"
			}
			name := pokerEntitys[i].Level.Name
			if len(name) == 1 {
				outputStr += name + " " + "|"
			} else {
				outputStr += name + "|"
			}
		}
		outputStr += "\n"
		for i := 0; i < len(pokerEntitys); i++ {
			if i == 0 {
				outputStr += "└──╯"
			} else {
				outputStr += "──╯"
			}
		}
	}
	fmt.Println(outputStr)
}

func textOnly(pokerEntitys []PokerEntity) {
	outputStr := ""
	if pokerEntitys != nil && len(pokerEntitys) > 0 {
		for i := 0; i < len(pokerEntitys); i++ {
			name := pokerEntitys[i].Level.Name
			pokerType := pokerEntitys[i].Type
			outputStr += name + pokerType
		}
	}
	fmt.Println(outputStr)
}

func textOnlyNoType(pokerEntitys []PokerEntity) {
	outputStr := ""
	if pokerEntitys != nil && len(pokerEntitys) > 0 {
		for i := 0; i < len(pokerEntitys); i++ {
			name := pokerEntitys[i].Level.Name
			outputStr += name + " "
		}
	}
	fmt.Println(outputStr)
}
