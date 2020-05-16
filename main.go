package main

import (
	"fmt"
	"go-ratel/network"
	"log"
	"net"
)

func main() {
	addr, e := net.ResolveTCPAddr("tcp", "39.105.65.8:1024")
	if e != nil {
		panic(e)
	}
	conn, e := net.DialTCP("tcp", nil, addr)
	if e != nil {
		panic(e)
	}

	for {
		codec := network.NewCodec(conn)
		e := codec.Read()
		if e != nil {
			log.Print(e)
		}
		for {
			transferData, b, e := codec.Decode()
			if e != nil {
				log.Panic(e)
			}
			if b {
				fmt.Println(transferData)
				continue
			}
			break
		}
	}
}
