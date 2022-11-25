package main

import "fmt"

func main() {
	fmt.Println("main called")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}
