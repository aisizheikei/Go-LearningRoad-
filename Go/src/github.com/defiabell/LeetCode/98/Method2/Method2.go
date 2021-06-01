//中序遍历实现
package Method2

import (
	"math"

	. "github.com/defiabell/LeetCode/98/TypeModel"
)

var lastValue = math.MinInt64

func IsValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !IsValidBST2(root.Left) {
		return false
	}
	if root.Val <= lastValue {
		return false
	}
	lastValue = root.Val
	return IsValidBST2(root.Right)
}
