/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：
输入：head = [1,2]
输出：[2,1]
示例 3：
输入：head = []
输出：[]
提示：链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/
package main

import "fmt"

func main() {
	args1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	printListNode(reverseList(args1))
}

func printListNode(head *ListNode) {
	if head != nil {
		for node := head; node != nil; node = node.Next {
			fmt.Println(node.Val)
		}
	}
}

//20210915
func reverseList_sys(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next //记录下个节点
		cur.Next = pre   //当前节点的next指向上一个节点
		pre = cur        //当前节点成为上个节点
		cur = temp       //下个节点成为当前节点
	}
	return pre //最后当前节点指向nil，上个节点指向最后一个节点

}

//----------------------------

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := reverseListRecursion(head, head.Next)
	head.Next = nil
	return result
}
func reverseListRecursion(node *ListNode, nextNode *ListNode) *ListNode {
	if nextNode == nil {
		return node
	}
	tempNode := nextNode.Next
	nextNode.Next = node
	return reverseListRecursion(nextNode, tempNode)
}
