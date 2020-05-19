package common

type PokerType = string
type PokerLevel = string

type Poker struct {
	Level PokerLevel
	Type  PokerType
}

type Level struct {
	Level int
	Name  string
	Alias []string
}

type Room struct {
	Id int
}

var pokerTypeMap = make(map[string]string)
var pokerLevelMap = make(map[string]Level)

func init() {
	pokerTypeMap["BLANK"] = " "
	pokerTypeMap["DIAMOND"] = "♦"
	pokerTypeMap["CLUB"] = "♣"
	pokerTypeMap["SPADE"] = "♠"
	pokerTypeMap["HEART"] = "♥"

	pokerLevelMap["LEVEL_3"] = Level{Level: 3, Name: "3", Alias: []string{"3"}}
	pokerLevelMap["LEVEL_4"] = Level{Level: 4, Name: "4", Alias: []string{"4"}}
	pokerLevelMap["LEVEL_5"] = Level{Level: 5, Name: "5", Alias: []string{"5"}}
	pokerLevelMap["LEVEL_6"] = Level{Level: 6, Name: "6", Alias: []string{"6"}}
	pokerLevelMap["LEVEL_7"] = Level{Level: 7, Name: "7", Alias: []string{"7"}}
	pokerLevelMap["LEVEL_8"] = Level{Level: 8, Name: "8", Alias: []string{"8"}}
	pokerLevelMap["LEVEL_9"] = Level{Level: 9, Name: "9", Alias: []string{"9"}}
	pokerLevelMap["LEVEL_10"] = Level{Level: 10, Name: "10", Alias: []string{"T", "t", "0"}}
	pokerLevelMap["LEVEL_J"] = Level{Level: 11, Name: "J", Alias: []string{"J", "j"}}
	pokerLevelMap["LEVEL_Q"] = Level{Level: 12, Name: "Q", Alias: []string{"Q", "q"}}
	pokerLevelMap["LEVEL_K"] = Level{Level: 13, Name: "K", Alias: []string{"K", "k"}}
	pokerLevelMap["LEVEL_A"] = Level{Level: 14, Name: "A", Alias: []string{"A", "a", "1"}}
	pokerLevelMap["LEVEL_2"] = Level{Level: 15, Name: "2", Alias: []string{"2"}}
	pokerLevelMap["LEVEL_SMALL_KING"] = Level{Level: 16, Name: "S", Alias: []string{"S", "s"}}
	pokerLevelMap["LEVEL_BIG_KING"] = Level{Level: 17, Name: "X", Alias: []string{"X", "x"}}
}

// 将Java枚举转换成struct实体
type PokerEntity struct {
	Type  PokerType
	Level Level
}

func ConvertEnumPoker(pokers *[]Poker) *[]PokerEntity {
	pokerEntitys := make([]PokerEntity, 0)
	for _, poker := range *pokers {
		pokerEntitys = append(pokerEntitys, PokerEntity{
			Type:  pokerTypeMap[poker.Type],
			Level: pokerLevelMap[poker.Level],
		})
	}
	return &pokerEntitys
}
