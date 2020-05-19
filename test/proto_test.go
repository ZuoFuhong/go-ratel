package test

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-ratel/common"
	"testing"
)

func Test_proto(t *testing.T) {
	clientData := common.ClientTransferDataProtoc{
		Code: "code",
		Data: "data",
		Info: "info",
	}

	encodeClientData, e := proto.Marshal(&clientData)
	if e != nil {
		panic(e)
	}

	clientData2 := common.ClientTransferDataProtoc{}
	e = proto.Unmarshal(encodeClientData, &clientData2)
	if e != nil {
		panic(e)
	}
	fmt.Println(clientData2)
}

func Test_Hex(t *testing.T) {
	fmt.Println(0xffffffff)
	fmt.Println(ComputeRawVarint32Size(127))
	fmt.Println(ComputeRawVarint32Size(128))
}

func ComputeRawVarint32Size(value int) int {
	if value&(0xffffffff<<7) == 0 {
		return 1
	}
	if value&(0xffffffff<<14) == 0 {
		return 2
	}
	if value&(0xffffffff<<21) == 0 {
		return 3
	}
	if value&(0xffffffff<<28) == 0 {
		return 4
	}
	return 5
}

func Test_BigEndian(t *testing.T) {
	x := uint16(255)

	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, x)
	fmt.Println(buf) // output: [0 255]
}

func Test_Varint(t *testing.T) {
	buf := make([]byte, 10)

	//x := uint64(127)
	x := uint64(128)
	tmp := proto.EncodeVarint(x)
	copy(buf, tmp)

	u, i := proto.DecodeVarint(buf)
	fmt.Println("u = ", u, " i = ", i)
}
