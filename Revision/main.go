package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

//client "client"
//	server "server"
//	"sync"

//httpserver "Revision/httpserver"
type ABC struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	/*var wg sync.WaitGroup
	wg.Add(2)
	go server.CreateServer(&wg)
	go client.CreateClient(&wg)
	wg.Wait()*/

	//httpserver.InitializeHttpServerMux()

	var sm sync.Map
	sm.Store("subham", ABC{10, "subham"})
	if res, ok := sm.Load("subham"); ok {
		fmt.Println("res : ", res)
	}

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(" key: ", key, " value: ", value)
		if key == "subham" {
			return false
		}
		return true

	})
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("open error : ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, dataLine := range lines {
		fmt.Println("line : ", dataLine)
	}
	defer file.Close()
	arr := []string{"asda", "sdad", "iieie", "uuada"}
	for dataLine := range arr {
		fmt.Println("line : ", dataLine)
	}
}
