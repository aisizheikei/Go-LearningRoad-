/*
描述
给定一个数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
示例1
输入：[2,3,4,5]
返回值：4
说明：[2,3,4,5]是最长子数组
示例2
输入：[2,2,3,4,3]
返回值：3
说明：[2,3,4]是最长子数组
示例3
输入：[9]
返回值：1
示例4
输入：[1,2,3,1,2,3,2,2]
返回值：3
说明：最长子数组为[1,2,3]
示例5
输入：[2,2,3,4,8,99,3]
返回值：5
说明：最长子数组为[2,3,4,8,99]
备注：
1 \leq n \leq 10^51≤n≤10^5
*/
package main

import "fmt"

func main() {
	//args1 := [...]int{1, 2, 3, 1, 2, 3, 2, 2}
	args2 := [...]int{3, 3, 2, 1, 3, 3, 3, 1}
	fmt.Println(maxLength(args2[:]))
}

var containMap map[int]int = make(map[int]int) //value,index
var headIndex int = 0                          //无重复数组起点索引
var maxLen int = 0                             //无重复数组最大长度

/**
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxLength(arr []int) int {
	// write code here
	if len(arr) == 0 {
		return 0
	}
	for i, v := range arr {
		val, ok := containMap[v]
		if ok { //前面已经遇到过了
			if headIndex <= val {
				headIndex = val + 1
			}
		}
		containMap[v] = i
		if i-headIndex+1 > maxLen {
			maxLen = i - headIndex + 1
		}
		//fmt.Println(containMap)
		//fmt.Println("maxLen:", maxLen, ",headIndex:", headIndex, ",i:", i)
		//fmt.Println("----------------------------------------------")
	}
	return maxLen
}
