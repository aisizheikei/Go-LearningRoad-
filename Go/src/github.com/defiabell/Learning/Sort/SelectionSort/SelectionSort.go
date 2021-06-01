//选择排序
package SelectionSort

func SelectionSort(arr []int) []int {
	//用最后一个数 和 前面的比较，最后最大值存到最后的位置。
	for i := len(arr) - 1; i > 0; i-- {
		index := i
		for j := 0; j < i; j++ {
			if arr[j] > arr[index] {
				index = j
			}
		}
		temp := arr[index]
		arr[index] = arr[i]
		arr[i] = temp
	}
	return arr
}
