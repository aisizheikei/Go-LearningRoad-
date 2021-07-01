/*
描述
输入一个链表，反转链表后，输出新链表的表头。
示例1
输入：
{1,2,3}
返回值：
{3,2,1}
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	args1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	result1 := ReverseList(args1)
	fmt.Println("result1:")
	for temp1, i := result1, 1; temp1 != nil; temp1, i = temp1.Next, i+1 {
		fmt.Printf("%v:%v\n", i, temp1.Val)
	}
	fmt.Println("----------------------------------------------")
}

func ReverseList_sys(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	var pre *ListNode = nil //当前的节点下一个，也就是新链表中当前节点的前一个   pre = cur
	cur := pHead            //当前节点     cur.Next = pre  ;cur = cur.Next

	for cur != nil {
		pre, cur.Next, cur = cur, pre, cur.Next
	}
	return pre
}
func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	// write code here
	result := new(ListNode)
	result.Val = pHead.Val
	tempNode := pHead.Next
	for ; tempNode != nil; tempNode = tempNode.Next {
		result = &ListNode{tempNode.Val, result}
		//fmt.Println(result)
	}
	return result
}
