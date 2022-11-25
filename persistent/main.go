package main

import "fmt"

type CharCount struct {
	Data string
}

func main() {
	// var a []byte
	// b := []int{1, 23, 13, 131}
	// for ind, _ := range b {
	// 	a = append(a, byte(b[ind]))
	// }
	// fmt.Println(string(a))

	a := [5]int{0, 1, 2, 3, 4}
	b := a[1:3]
	fmt.Println(a, b)
	b = append(b, 10)
	a = append(a, 10)
	fmt.Println(a, b)
}

func checkCount(obj *CharCount, ch chan string) {
	if obj == nil {
		ch <- ""
	} else {
		var res []rune
		mapData := make(map[rune]int)
		for _, char := range obj.Data {
			if val, ok := mapData[char]; !ok {
				mapData[char] = 1
			} else {
				mapData[char] = val + 1
			}
		}
		for key, mapVal := range mapData {
			res = append(res, key)
			valSt := fmt.Sprintf("%d", mapVal)
			res = append(res, []rune(valSt)...)
		}
		ch <- string(res)
	}
}
