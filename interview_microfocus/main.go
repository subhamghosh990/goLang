package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var i int
	if len(os.Args) > 0 {
		name = os.Args[0]
		for i = len(name) - 1; i >= 0; i-- {
			if name[i] == '\\' || name[i] == '/' {
				break
			}
		}
		tempByte := []byte(name)
		name = string(tempByte[i+1:])
	}
	fmt.Println("file name : ", name)
}
