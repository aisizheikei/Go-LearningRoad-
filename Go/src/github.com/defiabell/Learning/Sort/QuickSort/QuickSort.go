//快速排序
package QuickSort

func QuickSort(arr []int) []int {
	quickSort_recursion(arr, 0, len(arr)-1)
	return arr
}

func quickSort_recursion(arr []int, left, right int) {
	i, j := left, right
	//递归终止条件
	if right <= left {
		return
	}
	key := arr[i]
	for i < j {
		for i < j && arr[j] >= key {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= key {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = key
	quickSort_recursion(arr, left, i-1)
	quickSort_recursion(arr, i+1, right)
}
