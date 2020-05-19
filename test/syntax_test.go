package test

import (
	"encoding/json"
	"fmt"
	"go-ratel/common"
	"strconv"
	"testing"
)

func Test_Map(t *testing.T) {
	data := make(map[string]interface{})
	fmt.Println(data["name"]) // output: nil
	data["age"] = 24
	fmt.Println(data["age"].(string)) // occur error
}

func Test_Nexted(t *testing.T) {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "dazuo"
	dataMap["room"] = common.Room{Id: 1}
	bytes, _ := json.Marshal(&dataMap)

	dataMap2 := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &dataMap2)
	fmt.Print(dataMap2["room"].(map[string]interface{}))
}

func Test_Byte(t *testing.T) {
	bytes := []byte("hello world\t")
	for _, v := range bytes {
		if string(v) == "\t" {
			fmt.Println(v)
		}
	}
}

func Test_format(t *testing.T) {
	NICKNAME_MAX_LENGTH := 10
	base := "\n[%-4s] %-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s  surplus %-2s [%-8s]"
	sprintf := fmt.Sprintf(base, 123, "222", "333", "44")
	fmt.Print(sprintf)
}
