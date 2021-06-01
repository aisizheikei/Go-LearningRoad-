//算法
package main

import (
	"fmt"

	. "github.com/defiabell/Learning/Sort/BubbleSort"
	. "github.com/defiabell/Learning/Sort/QuickSort"
	. "github.com/defiabell/Learning/Sort/SelectionSort"
)

func main() {
	arr := []int{2, 3, 4, 6, 134, 76, 12, 98, 1, 1, 34, 9}
	//1.冒泡排序 BubbleSort
	//牛客题：失败了，时间超出
	arr_bubbleSort := make([]int, len(arr))
	copy(arr_bubbleSort, arr)
	bubbleSortResult := BubbleSort(arr_bubbleSort)
	fmt.Println("冒泡排序：", bubbleSortResult)
	//2.选择排序 SelectionSort
	arr_selectionSort := make([]int, len(arr))
	copy(arr_selectionSort, arr)
	selectionSortResult := SelectionSort(arr_selectionSort)
	fmt.Println("选择排序：", selectionSortResult)
	//3.快速排序 QuickSort
	arr_quickSort := make([]int, len(arr))
	copy(arr_quickSort, arr)
	quickSortResult := QuickSort(arr_quickSort)
	fmt.Println("快速排序", quickSortResult)
	//4.归并排序 MergeSOrt

	//5.堆排序 HeapSort

}
