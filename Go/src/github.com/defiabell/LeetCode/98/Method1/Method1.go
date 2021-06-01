package Method1

import . "github.com/defiabell/LeetCode/98/TypeModel"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsValidBST(root *TreeNode) bool {
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		} else if !IsValidBSTRecursion(root.Left, root.Val, root.Val) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		} else if !IsValidBSTRecursion(root.Right, root.Val, root.Val) {
			return false
		}
	}
	return true

}

func IsValidBSTRecursion(root *TreeNode, minVal, maxVal int) bool {
	compare1 := root.Val < minVal
	compare2 := root.Val < maxVal
	if root.Left != nil {
		if root.Left.Val == maxVal || root.Left.Val == minVal || compare2 == (root.Left.Val > maxVal) || compare1 == (root.Left.Val > minVal) || root.Left.Val >= root.Val {
			return false
		} else if !IsValidBSTRecursion(root.Left, minVal, root.Val) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val == maxVal || root.Right.Val == minVal || compare2 == (root.Right.Val > maxVal) || compare1 == (root.Right.Val > minVal) || root.Right.Val <= root.Val {
			return false
		} else if !IsValidBSTRecursion(root.Right, root.Val, maxVal) {
			return false
		}
	}
	return true

}
