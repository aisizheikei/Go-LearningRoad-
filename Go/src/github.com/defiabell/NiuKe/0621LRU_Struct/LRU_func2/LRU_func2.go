package LRU_func2

// key和value分别放在一个数组里，两个数组的索引是对应的。
// 第一个元素是最早的元素，当数组的数量等于k时，入队前先移除第一个元素，再把新元素加进来。
func LRU2(operators [][]int, k int) []int {
	result := make([]int, 0, len(operators))
	key := make([]int, k)   //存放key
	value := make([]int, k) //存放value

	for _, val := range operators {
		if val[0] == 1 { //插入
			//遍历key，查看是否存在该值，存在将该key放到数组第一个；不存在则判断数组大小，len==k时移除第一个并把该值加入数组。
			var index = -1
			for i, v := range key {
				if v == val[1] { //找到了该key
					key = append(key[1:], v)
					value = append(value[1:], val[2])
					index = i
					break
				}
			}
			if index == -1 { //没有找到key
				if len(key) == k { //数组到上限了
					//删除第一个，该值追加到最后一个
					key = append(key[1:], val[1])
					value = append(value[1:], val[2])
				} else {
					key = append(key, val[1])
					value = append(value, val[2])
				}
			}
		} else if val[0] == 2 { //获取
			var index = -1
			for i, v := range key {
				if v == val[1] { //找到该值了
					//把该值加入返回数组
					result = append(result, value[i])
					//把该值移动到最后一个
					key = append(key[:i], append(key[i+1:], v)...)
					value = append(value[:i], append(value[i+1:], value[i])...)
					index = i
					break
				}
			}
			if index == -1 { //没有找到该值
				//返回-1加入返回数组
				result = append(result, -1)
			}
		}
	}
	return result
}
