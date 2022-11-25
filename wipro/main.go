package main

import "fmt"

func main() {
	req := "aabb"
	res := removerChar(req)
	fmt.Println("res : ", res)
}

func removerChar(req string) string {
	var res []byte
	mapData := make(map[byte]bool)
	for i := 0; i < len(req); i++ {
		if _, ok := mapData[req[i]]; !ok {
			res = append(res, req[i])
			mapData[req[i]] = true
		}
	}
	return string(res)
}
