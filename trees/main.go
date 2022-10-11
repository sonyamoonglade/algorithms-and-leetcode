package main

import (
	"fmt"
)

func main() {

}

type Node struct {
	v     int
	left  *Node
	right *Node
}

func NewNode(left, right *Node, v int) *Node {
	return &Node{
		v:     v,
		left:  left,
		right: right,
	}
}

func dfsPrint(node *Node) {

	if node != nil {
		fmt.Println(node.v)
	}
	if node.left != nil {
		dfsPrint(node.left)
	}
	if node.right != nil {
		dfsPrint(node.right)
	}

}

func has(node *Node, dest int) bool {

	if node == nil {
		return false
	}

	if node != nil && node.v == dest {
		return true
	}

	return has(node.left, dest) || has(node.right, dest)
}

func treeSum(node *Node) int {
	if node == nil {
		return 0
	}

	return node.v + treeSum(node.left) + treeSum(node.right)
}

func treeHeight(node *Node) int {

	if node == nil {
		return 0
	}
	//assuming tree is perfect
	return treeHeight(node.right) + 1
}
