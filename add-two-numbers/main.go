package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	// fmt.Println("l1:", l1.Val, "l2:", l2.Val)

	sum := l1.Val + l2.Val
	node := ListNode{Val: sum}
	if sum > 9 {
		node.Val = node.Val % 10
		if l1.Next != nil {
			l1.Next.Val += 1
		} else if l2.Next != nil {
			l2.Next.Val += 1
		} else {
			l1.Next = &ListNode{Val: 1}
		}
	}
	node.Next = addTwoNumbers(l1.Next, l2.Next)
	return &node
}

func printResult(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println("")
}

func main() {
	l1 := ListNode{}
	l2 := ListNode{}
	result := &ListNode{}

	// l1 = [2,4,3], l2 = [5,6,4]
	// [7,0,8]
	l1 = ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 = ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result = addTwoNumbers(&l1, &l2)
	printResult(result)

	// l1 = [9, 9, 9, 9, 9, 9, 9], l2 = [9,9,9,9]
	// [8,9,9,9,0,0,0,1]
	l1 = ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 = ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}

	result = addTwoNumbers(&l1, &l2)
	printResult(result)

}
