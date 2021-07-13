/*
描述
请实现有重复数字的升序数组的二分查找
给定一个 元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1
示例1
输入：[1,2,4,4,5],4
返回值：2
说明：
从左到右，查找到第1个为4的，下标为2，返回2
示例2
输入：[1,2,4,4,5],3
返回值：-1
示例3
输入：[1,1,1,1,1],1
返回值：0
*/
package main

import "fmt"

func main() {
	args1 := []int{-2, 1, 2}
	fmt.Println(search(args1, -2))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 如果目标值存在返回下标，否则返回 -1
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
var index int = 0 //记录下标

func search(nums []int, target int) int {
	// write code here
	//1.如果数组为空，返回-1
	//2.取数组中位数比较，if(target==temp)，往前找，直到index==0或者temp<target
	//2.取数组中位数比较，if(target>temp)，后面的数据进行递归
	//3.if(target<temp)，前面的数据递归

	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] == target {
			return index
		} else {
			return -1
		}
	}
	lenth := len(nums)
	middle := lenth / 2
	if nums[middle] == target {
		for i := middle - 1; i >= 0; i-- {
			if i == 0 && nums[0] == target {
				return index
			}
			if nums[i] < target {
				index = index + i + 1
				return index
			}
		}
	} else if nums[middle] > target {
		return search(nums[0:middle], target)
	} else {
		index += middle
		return search(nums[middle:], target)
	}
	return -1
}
