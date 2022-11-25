package network

import (
	"fmt"
	"net"
	"os"
	"sync"
)

const (
	ADDR    = "localhost:990"
	NW_TYPE = "tcp"
)

var connServ net.Conn

func StartNetWorkSever(wg *sync.WaitGroup) {
	serv, err := net.Listen(NW_TYPE, ADDR)
	if wg != nil {
		defer wg.Done()
	}
	defer serv.Close()
	if err != nil {
		fmt.Println(" error while listening port ", err)
		os.Exit(1)
	}

	fmt.Println("waiting for client ")
	for {
		connServ, err := serv.Accept()
		if err != nil {
			fmt.Println(" error while accepting ", err)
		}
		res := readMessage(connServ)
		if res == true {
			break
		}

	}
}

func readMessage(connServ net.Conn) bool {
	buffer := make([]byte, 1024)
	for {
		mLen, err := connServ.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		} else {
			fmt.Println("Received: ", string(buffer[:mLen]))
			return true
		}
	}
}
func CloseServer() {
	if connServ != nil {
		defer connServ.Close()
	}
}
