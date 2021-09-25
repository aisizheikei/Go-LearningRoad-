/*easy
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：
输入：nums = [1]
输出：1
示例 3：
输入：nums = [0]
输出：0
示例 4：
输入：nums = [-1]
输出：-1
示例 5：
输入：nums = [-100000]
输出：-100000
提示：
1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105
链接：https://leetcode-cn.com/problems/maximum-subarray
*/
package main

import "fmt"

func main() {
	args1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} //
	fmt.Println(maxSubArray(args1))               //6
}

/*
思路：
维护两个变量：max--记录最大子序和；sum--记录遍历到当前节点的和
顺序遍历数组，如果sum>0,sum = sum+cur（当前节点）；else sum = cur
	max = Max（ sum，max）
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ { //从1开始，注意
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		max = getMax(max, sum)
	}
	return max
}
func getMax(first, second int) int {
	if first < second {
		return second
	}
	return first
}
