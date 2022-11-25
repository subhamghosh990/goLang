package main

import (
	"fmt"
	"strings"
	"sync"
)

func procudeData(ch chan int) {
	// if wg != nil {
	// 	defer wg.Done()
	// }
	for i := 1; i <= 10; i++ {
		ch <- 2 * i
	}
	close(ch)
}

func consumeData(wg *sync.WaitGroup, ch chan int) {
	if wg != nil {
		defer wg.Done()
	}
	for data := range ch {
		fmt.Println("Data received : ", data)
	}
}

func main() {
	// var wg sync.WaitGroup
	// ch := make(chan int)
	// wg.Add(5)
	// go procudeData(ch)
	// for i := 0; i < 5; i++ {
	// 	go consumeData(&wg, ch)
	// }
	// wg.Wait()

	data := []string{"flower", "flow", "flight"}
	res := findLongestCommonPreFix(data)
	fmt.Println(" Common preFix is ", res)

}

func findLongestCommonPreFix(arr []string) string {
	var res string
	if len(arr) == 0 {
		return res
	} else if len(arr) == 1 {
		return arr[0]
	}
	matchData := arr[0]
	mid := len(matchData) / 2
	var oneFound bool
	for {
		resp := loopData(mid, matchData, arr)
		fmt.Println("loopData resp : ", resp)
		if resp {
			mid += 1
			if mid >= len(matchData) {
				if oneFound == true {
					mid -= 1
					res = matchData[:mid]
					break
				}
			}
			oneFound = true
		} else {
			mid -= 1
			if oneFound == true {
				res = matchData[:mid]
				break
			}
		}
	}
	return res
}

func loopData(mid int, matchData string, arr []string) bool {
	midString := matchData[:mid]
	fmt.Println("mid : ", mid, " matchData : ", matchData, " midString : ", midString)
	for i := 1; i < len(arr); i++ {
		present := checkMidStringPresnt(midString, arr[i])
		if present == false {
			return false
		}

	}
	return true
}

func checkMidStringPresnt(midString string, data string) bool {
	if strings.HasPrefix(data, midString) {
		return true
	}
	return false
}
