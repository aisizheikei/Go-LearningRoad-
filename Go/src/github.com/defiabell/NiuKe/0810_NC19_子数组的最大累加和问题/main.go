/*
描述
给定一个数组arr，返回子数组的最大累加和
例如，arr = [1, -2, 3, 5, -2, 6, -1]，所有子数组中，[3, 5, -2, 6]可以累加出最大的和12，所以返回12.
题目保证没有全为负数的数据
[要求]
时间复杂度为O(n)O(n)，空间复杂度为O(1)O(1)

示例1
输入：[1, -2, 3, 5, -2, 6, -1]
返回值：12
*/
package main

import "fmt"

func main() {
	args1 := []int{1, -2, 3, 5, -2, 6, -1}
	fmt.Println(maxsumofSubarray(args1))
}

/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray(arr []int) int {
	// write code here
	max := 0
	temp := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]+temp > 0 {
			temp = arr[i] + temp
		} else {
			temp = 0
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
