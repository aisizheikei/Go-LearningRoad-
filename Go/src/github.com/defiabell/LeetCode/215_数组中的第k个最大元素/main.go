/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
提示：
1 <= k <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
*/

package main

import (
	"fmt"
	"math/rand"

	. "github.com/defiabell/LeetCode/215_数组中的第k个最大元素/findKthLargest2"
)

func main() {
	args1 := []int{3, 2, 1, 5, 6, 4}
	fmt.Println("func1:", findKthLargest(args1, 2))
	fmt.Println("func2:", FindKthLargest2(args1, 2))
}

func findKthLargest(nums []int, k int) int {
	var target = len(nums) - k
	left := 0
	right := len(nums) - 1
	for left < right {
		//使用随机数，因为用例可能是顺序数组，那样排序就变成了n^2，会超时
		index := rand.Intn(right-left) + left //注意right-left必须大于0
		temp := nums[index]
		nums[left], nums[index] = nums[index], nums[left]
		index = QuickSort(nums, left, right)
		if index == target {
			return temp
		} else if index < target { //目标值在较大的数组中
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return nums[left]
}

//快排，返回起始元素索引
func QuickSort(nums []int, left, right int) int {
	if left >= right {
		return -1
	}
	temp := nums[left]
	i, j := left, right
	for i < j {
		for ; i < j && nums[j] >= temp; j-- {

		}
		nums[i] = nums[j]
		for ; i < j && nums[i] <= temp; i++ {

		}
		nums[j] = nums[i]
	}
	nums[i] = temp
	//QuickSort(nums,left,i-1)
	//QuickSort(nums,i+1,right)

	return i

}
