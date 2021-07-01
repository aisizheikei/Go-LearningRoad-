package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	args1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	args1.Next.Next.Next = args1.Next
	result1 := hasCycle(args1)
	fmt.Println("result1:", result1)
}

/**
 *
 * @param head ListNode类
 * @return bool布尔型
 */
//该解使用一个字典，记录每个节点的地址，虽然能得到最终结果，但是不合题意，题目要求空间复杂度是O(1)。改解空间复杂度是O(n)
func hasCycle(head *ListNode) bool {
	// write code here
	nodeMap := make(map[*ListNode]bool, 100)
	if head == nil {
		return false
	}
	for tempNode := head; tempNode != nil; tempNode = tempNode.Next {
		_, ok := nodeMap[tempNode]
		if ok {
			return true
		}
		nodeMap[tempNode] = true
		//fmt.Println(nodeMap)
	}
	return false
}

//通过快慢指针，可以找到循环
func hasCycle_sys(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return false
	}
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		fast, slow = fast.Next.Next, slow.Next //一快一慢，如果是环，最多走 环前n+环m步后相遇
		if fast == slow {
			return true
		}
	}
}
