/*
大小堆方法解题
*/
package findKthLargest2

func FindKthLargest2(nums []int, k int) int {
	minHeap(nums, len(nums))
	//调整最大堆
	for i := 1; i < k; i++ {
		nums[len(nums)-i], nums[0] = nums[0], nums[len(nums)-i]
		minHeap(nums, len(nums)-i)
	}
	return nums[0]
}

//最大堆
func minHeap(nums []int, len int) {
	i := len/2 - 1      //从最大非叶子节点开始遍历
	var biggerIndex int //最大非叶子节点 左子节点
	for ; i >= 0; i-- {
		for j := i; j < len; {
			biggerIndex = 2*j + 1
			if biggerIndex >= len {
				break
			}
			if len > biggerIndex+1 && nums[biggerIndex] < nums[biggerIndex+1] {
				biggerIndex++ //指向又子节点
			}
			if nums[biggerIndex] > nums[j] {
				//如果子节点大，那么替换子节点与父节点
				nums[biggerIndex], nums[j] = nums[j], nums[biggerIndex]
				j = biggerIndex
			} else {
				break
			}
		}

	}
}
