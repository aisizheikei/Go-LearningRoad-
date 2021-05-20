package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 = &ListNode{2, &ListNode{Val: 3}}
	var l2 = &ListNode{2, &ListNode{Val: 5}}
	var l3 = mergeTwoLists(l1, l2)

	if l3 == nil {
		fmt.Println("集合为空")
	}
	fmt.Println("1", ":", l3.Val)
	i := 1
	temp := l3
	for {
		i++
		temp = temp.Next
		if temp == nil {
			return
		}
		fmt.Println(i, ":", temp.Val)
		if i > 100 {
			break
		}
	}
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var ret = new(ListNode)
	mergeLists(l1, l2, ret)
	return ret
}

func mergeLists(l1, l2 *ListNode, result *ListNode) {
	// if result == nil {
	// 	return
	// }
	if l1 == nil {
		*result = *l2
		return
	} else if l2 == nil {
		*result = *l1
		return
	} else if l1.Val > l2.Val {
		result.Val = l2.Val
		result.Next = new(ListNode)
		mergeLists(l1, l2.Next, result.Next)
		return
	} else {
		result.Val = l1.Val
		result.Next = new(ListNode)
		mergeLists(l1.Next, l2, result.Next)
		return
	}
}
