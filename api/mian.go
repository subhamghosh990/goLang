package main

import "fmt"

func main() {
	// n := 5
	// temp := n
	// for temp >= 0 {
	// 	fmt.Print(temp)
	// 	temp--
	// }
	// temp = 1
	// for temp <= n {
	// 	fmt.Print(temp)
	// 	temp++
	// }
	arr, i := []int{8, 9, 9, 0, 7, 9, 0}, 0
	fmt.Println(arr)
	arrMap := make(map[int]struct{})
	for i < len(arr) {
		if _, ok := arrMap[arr[i]]; ok {
			arr = append(arr[:i], arr[i+1:]...)
			continue
		} else {
			arrMap[arr[i]] = struct{}{}
		}
		i++
	}
	fmt.Println(arr)
}
