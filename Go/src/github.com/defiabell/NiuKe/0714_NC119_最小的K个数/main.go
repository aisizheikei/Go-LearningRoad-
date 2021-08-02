/*
描述
给定一个数组，找出其中最小的K个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。
0 <= k <= input.length <= 10000
0 <= input[i] <= 10000

示例1
输入：[4,5,1,6,2,7,3,8],4
返回值：[1,2,3,4]
说明：返回最小的4个数即可，返回[1,3,2,4]也可以
示例2
输入：[1],0
返回值：[]
示例3
输入：[0,1,2,1,2],3
返回值：[0,1,1]

*/

package main

import "fmt"

func main() {
	args1 := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers_Solution(args1, 4))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	//1.维护一个k大小的数组。每次比较一个值与数组中数据，取出其中最大的删除。最后剩下的数组元素返回
	var minList []int = make([]int, 0)
	for i := 0; i < len(input); i++ {
		handlerMinList(&minList, k, input[i])
	}
	return minList
}

func handlerMinList(minList *[]int, k, val int) { //注意要传指针，虽然slice是引用类型，但是slice和map一样当参数传入的时候会带上长度信息，所以如果不传指针，最后返回的数据有可能不全（长度等同传入时）
	if minList == nil {
		return
	}
	if len(*minList) < k { //不足指定大小的数组就加入
		//fmt.Println("if:", *minList)
		*minList = append(*minList, val)
		//fmt.Println("if:", *minList)

	} else { //超过指定大小，置换其中最大的一个
		for i := 0; i < len(*minList); i++ {
			if (*minList)[i] > val {
				(*minList)[i], val = val, (*minList)[i]
			}
		}
		//fmt.Println("else:", *minList)
	}
}
