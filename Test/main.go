package main

import "fmt"

func main() {
	row := 4
	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
