package network

import (
	"go-ratel/common"
	"log"
	"net"
	"time"
)

// TCP连接上下文（网络层）
type ConnContext struct {
	codec      *Codec
	clientChan *chan common.ClientTransferDataProtoc
	serverChan *chan common.ServerTransferDataProtoc
}

func NewConnContext(conn *net.TCPConn, clientChan *chan common.ClientTransferDataProtoc, serverChan *chan common.ServerTransferDataProtoc) *ConnContext {
	return &ConnContext{
		codec:      NewCodec(conn),
		clientChan: clientChan,
		serverChan: serverChan,
	}
}

func (ctx *ConnContext) DoConn() {
	ctx.DoSend()
	for {
		e := ctx.codec.Read()
		if e != nil {
			log.Print(e)
		}
		for {
			transferData, b, e := ctx.codec.Decode()
			if e != nil {
				log.Panic(e)
			}
			if b {
				*ctx.clientChan <- *transferData
				continue
			}
			break
		}
	}
}

func (ctx *ConnContext) DoSend() {
	go func() {
		for {
			transferData := <-*ctx.serverChan
			err := ctx.codec.Encode(&transferData, time.Second*10)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
