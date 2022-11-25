package main

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

type Tree struct {
	Head *Node
}

    //     10
    // -19     20							-  1. lVal = -19, 2. lval = 15, 2.Rval = 7
    //        15   7
func maxPath(head *Node, value *int, prevHead *Node) int {

	if head == nil {
		return 0
	} else if head.Left == nil && head.Right == nil {
		return head.Val
	}
	lVal := maxPath(head.Left, value, head)
	rVal := maxPath(head.Right, value, head)
	if rVal > lVal {
		*value = *value + rVal
	} else {
		*value = *value + lVal
	}
	return 0
}


func checkIsland([][]int, Row, col int) int {
	res := 0
	for i:= 0, i < Row, i++ {
		for j:=0 , j < col ; j++ {
			if[]
		}
	}
}
func main() {
	tree := Tree{}
	var res int
	maxPath(tree.Head, &res)
	if tree.Head != nil {
		res = res + tree.Head.Val
	}
}
