package main

import (
	"testing"
	"fmt"
)

func TestcheckCount() {
	obj := CharCount{Data: "aabbcccddd"}
	ch := make(chan string)
	go checkCount(&obj, ch)
	res := <-ch
	fmt.Println("response received : ", res)
}
