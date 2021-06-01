package main

import (
	"fmt"

	. "github.com/defiabell/LeetCode/98/Method1" //. 使用该包内的接口可以不用带命名空间
	. "github.com/defiabell/LeetCode/98/Method2" //. 使用该包内的接口可以不用带命名空间
	. "github.com/defiabell/LeetCode/98/TypeModel"
)

func main() {
	tree1 := TreeNode{
		Val:   5,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
	}
	tree2 := TreeNode{
		0, nil, nil,
	}
	result := IsValidBST(&tree1)
	fmt.Println("method1--tree1结果是：", result)
	//fmt.Println("method1--tree1结果是：", IsValidBST2(&tree1))
	fmt.Println("method1--tree2结果是：", IsValidBST2(&tree2))
}
