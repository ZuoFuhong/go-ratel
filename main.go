package main

import (
	"go-ratel/common"
	"go-ratel/event"
	"go-ratel/network"
	"log"
	"net"
)

func main() {
	start("39.105.65.8:1024")
}

func start(address string) {
	addr, e := net.ResolveTCPAddr("tcp", address)
	if e != nil {
		log.Panic(e)
	}
	conn, e := net.DialTCP("tcp", nil, addr)
	if e != nil {
		log.Panic(e)
	}

	// 用两个chan用于 应用层 与 网络层 之间的解耦
	clientChan := make(chan common.ClientTransferDataProtoc)
	serverChan := make(chan common.ServerTransferDataProtoc)

	ectx := event.NewEventContext(&clientChan, &serverChan)
	ectx.DoListen()

	cctx := network.NewConnContext(conn, &clientChan, &serverChan)
	cctx.DoConn()
}
