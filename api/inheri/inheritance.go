package main

import (
	"fmt"
	"sync"
)

type Base struct {
	Id int
}

type Derived struct {
	Base
	Name string
}

func main() {
	// obj := Derived{}
	// obj.Name = "Subham"
	// obj.Id = 990

	// fmt.Println(obj.Name, obj.Id)
	// wg := sync.WaitGroup{}
	// wg.Add(2)
	// go Test1(&wg)
	// go Test2(&wg)
	// wg.Wait()
	var n interface{}
	n = -5
	switch n.(type) {
	case int:
		fmt.Println("input type is expected")
		if v, ok := n.(int); ok {
			printRecursive(v)
		}

	default:
		fmt.Println("input type is not expected\nPlease provide int type")

	}

}
func printRecursive(n int) {
	if n < 0 {
		return
	}
	fmt.Println(n)
	printRecursive(n - 1)
}
func Test1(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	n := 10
	for n > 0 {
		fmt.Println("test1")
		n--
	}
}
func Test2(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	n := 10
	for n > 0 {
		fmt.Println("test2")
		n--
	}
}
