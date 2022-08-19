package main

import (
	"fmt"
	"sync"
)

func sendChannel(ch chan int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	fmt.Println("sending 1 to rcvChanel")
	ch <- 1
}

func rcvChanel(ch chan int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	select {
	case val := <-ch:
		fmt.Println("recieved val is : ", val)
	}

}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go sendChannel(ch, &wg)
	go rcvChanel(ch, &wg)
	wg.Wait()
}
