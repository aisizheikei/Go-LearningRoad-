/*
描述
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
注：本题保证二叉树中每个节点的val值均不相同。
示例1
输入：[3,5,1,6,2,0,8,#,#,7,4],5,1
返回值：3
*/

package main

import "fmt"

//[3,5,1,6,2,0,8,#,#,7,4]
func main() {
	args1 := TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}
	fmt.Println(lowestCommonAncestor(&args1, 5, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var parentNodeMap map[int]int = make(map[int]int, 0) //记录每个节点的上级节点数字
/**
 *
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	getMapOfNode(root)
	var o1ParentSet map[int]bool = make(map[int]bool, 0) //o1的上级
	o1ParentSet[o1] = true
	temp1 := o1
	for {
		val, ok := parentNodeMap[temp1]
		if ok {
			o1ParentSet[val] = true
			temp1 = val
		} else {
			break
		}
	}
	temp2 := o2
	for { //遍历o2上级，直到遇到o1的上级
		_, ok := o1ParentSet[temp2]
		if ok { //遇到共同祖先了
			return temp2
		}
		temp2 = parentNodeMap[temp2]
	}
}
func getMapOfNode(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		parentNodeMap[node.Left.Val] = node.Val
		getMapOfNode(node.Left)
	}
	if node.Right != nil {
		parentNodeMap[node.Right.Val] = node.Val
		getMapOfNode(node.Right)
	}
	//fmt.Println(parentNodeMap)
}

// 最近公共祖先和o1,o2有三种关系：
// o1,o2分别在祖先左右两侧
// 祖先是o1，o2在祖先左/右侧
// 祖先是o2，o1在祖先左/右侧
func lowestCommonAncestor_sys(root *TreeNode, o1 int, o2 int) int {
	// write code here
	return CommonAncestor(root, o1, o2).Val
}

func CommonAncestor(root *TreeNode, o1, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	left := CommonAncestor(root.Left, o1, o2)
	right := CommonAncestor(root.Right, o1, o2)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
