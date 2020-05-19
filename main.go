package main

import (
	"bytes"
	"encoding/json"
	"go-ratel/command"
	"go-ratel/common"
	"go-ratel/event"
	"go-ratel/network"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	host, port := parseFlag()
	if port == 0 {
		host, port = selectServer()
	}
	start(host, port)
}

func parseFlag() (string, int) {
	if len(os.Args) == 5 {
		host := os.Args[2]
		port, e := strconv.Atoi(os.Args[4])
		if e == nil {
			return host, port
		}
	}
	return "", 0
}

func selectServer() (string, int) {
	servers := getServerList()
	serverList := make([]string, 0)
	e := json.Unmarshal([]byte(servers), &serverList)
	if e != nil {
		log.Panic(e)
	}
	command.PrintNotice("Please select a server:")
	for i := 0; i < len(serverList); i++ {
		command.PrintNotice(strconv.Itoa(i+1) + ". " + serverList[i])
	}
	serverPick := -1
	for {
		serverPick, e = strconv.Atoi(command.Write("option"))
		if e == nil && serverPick > 0 {
			if serverPick > len(serverList) {
				command.PrintNotice("The server address does not exist!")
				continue
			}
			break
		}
	}
	serverAddress := strings.Split(serverList[serverPick-1], ":")
	host := serverAddress[0]
	port, _ := strconv.Atoi(serverAddress[1])
	return host, port
}

func getServerList() string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://raw.githubusercontent.com/ainilili/ratel/master/serverlist.json")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

func start(host string, port int) {
	addr, e := net.ResolveTCPAddr("tcp", host+":"+strconv.Itoa(port))
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
