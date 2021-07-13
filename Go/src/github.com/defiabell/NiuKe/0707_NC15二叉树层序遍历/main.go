/*
描述
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},

该二叉树层序遍历的结果是
[
[3],
[9,20],
[15,7]
]

示例1
输入：{1,2}
返回值：[[1],[2]]
示例2
输入：{1,2,3,4,#,#,5}
返回值：[[1],[2,3],[4,5]]

*/

package main

import "fmt"

func main() {
	args1 := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(levelOrder(&args1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here
	var result [][]int = make([][]int, 0)           //记录结果
	var nextNode []*TreeNode = make([]*TreeNode, 0) //记录下一层的node节点
	var lastNode []*TreeNode = make([]*TreeNode, 0) //记录下一层的下一层node节点
	if root == nil {
		return nil
	}
	result = append(result, []int{root.Val})
	nextNode = addNextNode(root, nextNode)
	for len(nextNode) > 0 {
		tempArr := make([]int, 0) //记录当前层数据
		for i := 0; i < len(nextNode); i++ {
			tempArr = append(tempArr, nextNode[i].Val) //记录当前节点数据
			lastNode = addNextNode(nextNode[i], lastNode)
		}
		//将当前层数据tempArr加入到result中
		result = append(result, tempArr)
		//重置nextNode和lastNode
		nextNode = lastNode
		lastNode = make([]*TreeNode, 0)
	}
	return result
}

//处理当前节点加入到节点队列中
func addNextNode(node *TreeNode, nodeList []*TreeNode) []*TreeNode {
	if node.Left != nil {
		nodeList = append(nodeList, node.Left) //左节点加进去
	}
	if node.Right != nil {
		nodeList = append(nodeList, node.Right) //右节点加进去
	}
	return nodeList
}
