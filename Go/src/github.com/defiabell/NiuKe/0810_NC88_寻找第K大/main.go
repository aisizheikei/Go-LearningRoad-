/*
描述
有一个整数数组，请你根据快速排序的思路，找出数组中第 大的数。

给定一个整数数组 ,同时给定它的大小n和要找的 ，请返回第 大的数(包括重复的元素，不用去重)，保证答案存在。
要求时间复杂度
示例1
输入：[1,3,5,2,2],5,3
返回值：2
示例2
输入：[10,10,9,9,8,7,5,6,4,3,4,2],12,3
返回值：9
说明：去重后的第3大是8，但本题要求包含重复的元素，不用去重，所以输出9
*/

package main

import "fmt"

func main() {
	args1 := []int{1, 3, 5, 2, 2}
	fmt.Println(findKth(args1, 5, 3))
}

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// write code here
	return quickSort(a)[n-K]
}
func quickSort(a []int) (result []int) {
	if len(a) <= 1 {
		result = a
		return
	}
	tempValue := a[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(a); i++ {
		if a[i] > tempValue {
			right = append(right, a[i])
		} else {
			left = append(left, a[i])
		}
	}
	if len(left) > 0 {
		result = append(result, quickSort(left)...)
	}
	result = append(result, tempValue)
	if len(right) > 0 {
		result = append(result, quickSort(right)...)
	}
	return
}
