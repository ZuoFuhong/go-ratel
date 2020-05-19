package test

import (
	"encoding/json"
	"fmt"
	"go-ratel/command"
	"go-ratel/common"
	"io/ioutil"
	"os"
	"testing"
)

func Test_PrintPokers(t *testing.T) {
	//command.PrintPokers()
	fmt.Println(os.Getwd())
	bytes, e := ioutil.ReadFile("./show_pokers.json")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes))
	dataMap := make(map[string]interface{})
	e = json.Unmarshal(bytes, &dataMap)
	if e != nil {
		panic(e)
	}
	// 需要两次转换
	pokers, e := json.Marshal(dataMap["pokers"])
	if e != nil {
		panic(e)
	}
	pokerList := make([]common.Poker, 0)
	e = json.Unmarshal([]byte(pokers), &pokerList)
	if e != nil {
		panic(e)
	}
	command.PrintPokers(pokerList, 0)
}
