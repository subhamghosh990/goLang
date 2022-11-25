package main

import "fmt"

type queueMsg struct {
	Data string
	Next *queueMsg
}

type node struct {
	HEAD *queueMsg
}

func (n *node) Push(_data string) {
	if n.HEAD == nil {
		n.HEAD = new(queueMsg)
		n.HEAD.Data = _data
		n.HEAD.Next = nil
	} else {
		temp := n.HEAD
		for temp.Next != nil {
			temp = temp.Next
		}
		newObj := new(queueMsg)
		newObj.Data = _data
		newObj.Next = nil
		temp.Next = newObj
	}
}

func (n *node) Pop() string {
	var res string
	if n.HEAD != nil {
		res = n.HEAD.Data
		n.HEAD = n.HEAD.Next
	}
	return res
}

func main() {
	var m map[string]int
	fmt.Println("m[\"A\"] ", m["A"])
}
