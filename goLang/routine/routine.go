package routine

import (
	"fmt"
	"sync"
)

func NormalGoRoutine() {
	for i := 0; i < 10; i++ {
		fmt.Println("normalGoRoutine : i : ", i)
	}
}

func WaitGroupGoRoutine(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("WaitGroupGoRoutine : i : ", i)
	}
}

func NormalChanGoRoutineString(ch chan string) {
	fmt.Println("NormalChanGoRoutine respoding Hello")
	ch <- "HELLO"
	close(ch)
}

func SendChanGoRoutine(ch chan<- int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("sendChanGoRoutine : i : ", i)
		ch <- i
	}
	close(ch)
}
func ReceivedChanGoRoutine(ch <-chan int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	isClosed := false
	for {
		select {
		case resp, ok := <-ch:
			if ok {
				fmt.Println("receivedChanGoRoutine : resp : ", resp)
			} else {
				isClosed = true
			}
		}
		if isClosed {
			break
		}
	}
	fmt.Println("receivedChanGoRoutine looping 100000")
	for i := 0; i < 100000; i++ {
	}
	fmt.Println("receivedChanGoRoutine looping 100000 end")
}
