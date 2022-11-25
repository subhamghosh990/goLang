package network

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var connClient net.Conn

func StartNetWorkClient(wg *sync.WaitGroup) {
	connClient, err := net.Dial(NW_TYPE, ADDR)
	if wg != nil {
		defer wg.Done()
	}
	if err != nil {
		fmt.Println(" error while listening port ", err)
		os.Exit(1)
	}
	go writeMessage(connClient)
}

func writeMessage(connClient net.Conn) {

	connClient.Write([]byte("Hello Server! Greetings."))
	fmt.Println(" client sent ")
}

func CloseClient() {
	if connClient != nil {
		defer connClient.Close()
	}

}
