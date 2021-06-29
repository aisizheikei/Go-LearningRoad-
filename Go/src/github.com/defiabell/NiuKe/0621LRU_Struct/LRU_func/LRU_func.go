//结果不是很对，暂不知那步出错。

package LRU_func

/*
描述:
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值

[要求]
set和get方法的时间复杂度为O(1)
某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案

示例1
输入：
[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
返回值：
[1,-1]
说明：
第一次操作后：最常使用的记录为("1", 1)
第二次操作后：最常使用的记录为("2", 2)，("1", 1)变为最不常用的
第三次操作后：最常使用的记录为("3", 2)，("1", 1)还是最不常用的
第四次操作后：最常用的记录为("1", 1)，("2", 2)变为最不常用的
第五次操作后：大小超过了3，所以移除此时最不常使用的记录("2", 2)，加入记录("4", 4)，并且为最常使用的记录，然后("3", 2)变为最不常使用的记录
*/

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

type Node struct {
	next *Node
	pre  *Node
	data int
}

var head = Node{new(Node), nil, 0}
var last = Node{nil, &head, 0}
var lruMap map[int]*Node

var result = make([]int, 0, 3)

var maxlen int

func LRU(operators [][]int, k int) []int {
	//初始化head和last
	head.next = &last
	// write code here
	//0 根据k生成一个k位的map
	lruMap = make(map[int]*Node, k) //不能再声明容量。map特殊
	maxlen = k
	//1	遍历operators
	for _, val := range operators {
		switch val[0] {
		case 1: //set
			setLRUMap(val[1], val[2])
		case 2: //get
			getval := getLRUMap(val[1])
			result = append(result, getval)
		}
	}
	return result
}

//根据正确答案推测题目的意思跟理解的有偏差。set的时候，如果存在该值，也会删除最后一个，把新值加进去。这就不善了。
//经过验证发现也不是这样，报错还是照样...
func setLRUMap_niuke(k, v int) {
	//不存在
	newNode := Node{head.next, &head, v}
	//判断是否超出限制了
	if len(lruMap) >= maxlen {
		//fmt.Println("remove last:", last.pre)
		//移除map中的数据
		delete(lruMap, last.pre.data)
		//移除最后一个，新值添加到第一个
		last.pre = last.pre.pre
		last.pre.next = &last

		head.next.pre = &newNode
		head.next = &newNode
	} else {
		//fmt.Println("newNode", newNode)
		head.next.pre = &newNode
		head.next = &newNode
	}
	//lruMap中加入新值
	lruMap[k] = &newNode
	//fmt.Println("head.next:", head.next)
}

func setLRUMap(k, v int) {
	val, ok := lruMap[k]
	if ok { //已经存在
		val.data = v
		//fmt.Println("V:", v)
		val.next.pre = val.pre
		val.pre.next = val.next
		head.next.pre = val
		val.next = head.next
		val.pre = &head
		head.next = val
	} else { //不存在
		newNode := Node{head.next, &head, v}
		//判断是否超出限制了
		if len(lruMap) >= maxlen {
			//fmt.Println("remove last:", last.pre)
			//移除map中的数据
			delete(lruMap, last.pre.data)
			//移除最后一个，新值添加到第一个
			last.pre = last.pre.pre
			last.pre.next = &last

			head.next.pre = &newNode
			head.next = &newNode
		} else {
			//fmt.Println("newNode", newNode)
			head.next.pre = &newNode
			head.next = &newNode
		}
		//lruMap中加入新值
		lruMap[k] = &newNode
		//fmt.Println("head.next:", head.next)
	}

}

func getLRUMap(k int) int {
	val, ok := lruMap[k]
	if ok { //已经存在
		val.next.pre = val.pre
		val.pre.next = val.next
		head.next.pre = val
		val.next = head.next
		val.pre = &head
		head.next = val
		return val.data
	} else {
		return -1
	}
}
