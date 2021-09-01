/*
示例1
输入：{1,2,3}
返回值：[[1,2,3],[2,1,3],[2,3,1]]
备注：n <= 10^6
*/

package main

import "fmt"

func main() {
	args1 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(threeOrders(&args1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here
	result := make([][]int, 0)
	result = append(result, firstOrder(root))
	result = append(result, middleOrder(root))
	result = append(result, lastOrder(root))
	return result
}

func firstOrder(root *TreeNode) (list []int) {
	if root != nil {
		list = append(list, root.Val)
		list = append(list, firstOrder(root.Left)...)
		list = append(list, firstOrder(root.Right)...)
	}
	return list
}
func middleOrder(root *TreeNode) (list []int) {
	if root != nil {
		list = append(list, middleOrder(root.Left)...)
		list = append(list, root.Val)
		list = append(list, middleOrder(root.Right)...)
	}
	return list
}
func lastOrder(root *TreeNode) (list []int) {
	if root != nil {
		list = append(list, lastOrder(root.Left)...)
		list = append(list, lastOrder(root.Right)...)
		list = append(list, root.Val)
	}
	return list
}
