package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(d int, ch chan int, wg *sync.WaitGroup) {
			if wg != nil {
				defer wg.Done()
			}
			if d%2 == 1 {
				ch <- d
			}
		}(v, ch, &wg)
	}
	wg.Wait()
	close(ch)
	for d := range ch {
		fmt.Println("received odd number is : ", d)
	}

}

func receivedData(ch chan int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	var close bool
	for {
		select {
		case isOdd, close := <-ch:
			if !close {
				fmt.Println("received value : ", isOdd)
			}
		}
		if close {
			fmt.Println("channel closed")
			break
		}
	}
}
