package main

import (
	"math/big"
	"strconv"
	"strings"
)

func main() {
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	s3 := &ListNode{
		Val:  6,
		Next: nil,
	}

	s2 := &ListNode{
		Val:  5,
		Next: s3,
	}

	s1 := &ListNode{
		Val:  3,
		Next: s2,
	}

	addTwoNumbers(l1, s1)
}

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	v1 := Traverse(l1)
	v2 := Traverse(l2)

	sum := Sum(v1, v2)

	return Listify(sum)
}

func Listify(ss string) *ListNode {

	var head *ListNode
	for _, s := range strings.Split(ss, "") {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		l := &ListNode{
			Val:  v,
			Next: head,
		}
		head = l
	}

	return head
}

func Sum(s1, s2 string) string {
	var v1, v2 = new(big.Int), new(big.Int)
	_, ok := v1.SetString(s1, 10)
	if !ok {
		panic("could not set string")
	}
	_, ok = v2.SetString(s2, 10)
	if !ok {
		panic("could no set secnod string")
	}

	v1.Add(v1, v2)

	return v1.String()
}

func Traverse(l *ListNode) string {
	if l == nil {
		return ""
	}
	return Traverse(l.Next) + strconv.Itoa(l.Val)
}
