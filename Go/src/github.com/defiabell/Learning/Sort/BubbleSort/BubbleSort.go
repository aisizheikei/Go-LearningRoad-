//冒泡排序
package Bubblesort

func BubbleSort(arr []int) []int {
	//1.遍历，比较两个数大小，大的放后面。一遍后最大的数就在最后了
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		// for index, _ := range arr {
		// 	if index > i-1 {
		// 		break
		// 	}
		// 	if arr[index] > arr[index+1] {
		// 		temp := arr[index+1]
		// 		arr[index+1] = arr[index]
		// 		arr[index] = temp
		// 	}
		// }
	}
	return arr
}
