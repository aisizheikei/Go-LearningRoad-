/*
描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
示例1
输入：2
返回值：2
示例2
输入：7
返回值：21
*/

package main

import "fmt"

func main() {
	fmt.Println("mine:", jumpFloor(5))
	fmt.Println("sys:", jumpFloor_sys(5))
	fmt.Println("super:", jumpFloor_super(5))
}

var floorCountMap map[int]int = make(map[int]int, 0) //记录到达指定台阶的步数。优化减少计算量

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	//1、jumpFloor(n) = jumpFloor(n-1) + jumFloor(n-2) 此为递归。
	//2、优化：递归会存在大量重复计算。例如jumpFloor(5) 就会计算两遍jumpFloor(3)。所以可以 动态规划优化
	if number <= 0 {
		return 0
	} else if number == 1 {
		return 1
	} else if number == 2 {
		return 2
	} else {
		val1, ok := floorCountMap[number-1]
		if !ok {
			val1 = jumpFloor(number - 1)
			//floorCountMap[number-1] = val1
		}
		val2, ok := floorCountMap[number-2]
		if !ok {
			val2 = jumpFloor(number - 2)
			//floorCountMap[number-2] = val2
		}
		floorCountMap[number] = val1 + val2
		//fmt.Println(floorCountMap)
		return floorCountMap[number]
	}
}

//正常解
func jumpFloor_sys(number int) int {
	dp := make([]int, number+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

//超级解
//空间复杂度为O(1)
//发现计算f[5]的时候只用到了f[4]和f[3], 没有用到f[2]...f[0],所以保存f[2]..f[0]是浪费了空间。
func jumpFloor_super(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	a, b := 1, 1
	c := 0
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
