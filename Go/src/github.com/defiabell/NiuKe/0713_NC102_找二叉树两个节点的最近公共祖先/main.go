/*
描述
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
注：本题保证二叉树中每个节点的val值均不相同。
示例1
输入：[3,5,1,6,2,0,8,#,#,7,4],5,1
返回值：3
*/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here

}
