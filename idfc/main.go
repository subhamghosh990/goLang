package main

type Node struct {
	Id          int
	Left, Right *Node
}

type Tree struct {
	Head *Node
}

//		1
//	2      3
//		1
//	3	 	2
func InvertTree(head *Node) {
	if head == nil {
		return
	} else {
		var temp *Node
		InvertTree(head.Left)
		InvertTree(head.Right)
		temp = head.Left
		head.Left = head.Right
		head.Right = temp
	}
}

func main() {
	tree := Tree{}
	InvertTree(tree.Head)
}
