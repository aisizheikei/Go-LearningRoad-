/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。
示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
提示：
树中节点数目范围是 [1, 3 * 10^4]
-1000 <= Node.val <= 1000
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	args1 := &TreeNode{
		-10,
		&TreeNode{9, &TreeNode{21, nil, nil}, &TreeNode{19, nil, nil}},
		&TreeNode{20, &TreeNode{25, nil, nil}, &TreeNode{8, nil, nil}},
	}
	fmt.Println(maxPathSum(args1))
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1.递归解法：
需要维护两个变量：
	一个全局最大值maxPath，
	一个经过左子节点的贡献最大值maxLeft(小于0的话为0)
	一个经过右子节点的贡献最大值maxRight(小于0的话为0)
递归返回值为该节点贡献最大值（val+max(maxLeft,maxRight)）
maxPath = max(maxPath, maxLeft + maxRight + val)
*/
var maxPath = math.MinInt32

func maxPathSum(root *TreeNode) int {
	maxPath = math.MinInt32 //因为是个全局变量，慎用，有风险，所以在用之前初始化一下。但是并发的时候就会有问题了
	getMaxPathSum(root)
	return maxPath
	//return 0
}
func getMaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := max(getMaxPathSum(root.Left), 0)
	maxRight := max(getMaxPathSum(root.Right), 0)
	maxPath = max(maxPath, maxLeft+maxRight+root.Val)
	return root.Val + max(maxLeft, maxRight)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
