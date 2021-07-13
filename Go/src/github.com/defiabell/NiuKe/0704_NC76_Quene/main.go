/*
描述:
用两个栈来实现一个队列，分别完成在队列尾部插入整数(push)和在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

示例:
输入:["PSH1","PSH2","POP","POP"]
返回:1,2
解析:
"PSH1":代表将1插入队列尾部
"PSH2":代表将2插入队列尾部
"POP“:代表删除一个元素，先进先出=>返回1
"POP“:代表删除一个元素，先进先出=>返回2
示例1
输入：["PSH1","PSH2","POP","POP"]
返回值：1,2
*/

package main

import "fmt"

func main() {
	Push(1)
	Push(2)
	fmt.Println(Pop())
	fmt.Println(Pop())
}

var stack1 []int //用于push，这样数据进来就是一个入队。
var stack2 []int //要出队的时候，将stack1的数据从后往前一个个添加到stack2.这样元素顺序就反了个，最后一个元素就是最早的数据。

func Push(node int) {
	//判断stack2是否有数据，如果有数据先转移到stack1，然后清空stack2
	//将新元素加到stack1
	if stack1 == nil {
		initStack1()
	}
	if stack2 == nil {
		initStack2()
	}
	if len(stack2) > 0 {
		//stack2数据加到stack1
		initStack1()
		for i := len(stack2) - 1; i >= 0; i-- {
			stack1 = append(stack1, stack2[i])
		}
		//清空stack2
		initStack2()
	}
	stack1 = append(stack1, node)

}

func Pop() int {
	//判断stack1是否有数据，如果有数据先转移到stack2，然后清空stack1
	//判断stacke2是否有数据，有则删除stack2最后一个元素，没有返回？？？
	if stack1 == nil {
		initStack1()
	}
	if stack2 == nil {
		initStack2()
	}
	if len(stack1) > 0 {
		//转移到stack2
		initStack2()
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		//清空stack1
		initStack1()
	}
	if len(stack2) == 0 {
		//???
		return -1
	}
	//取出最后一个并删除
	result := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return result
}

func initStack1() {
	stack1 = make([]int, 0, 20)
}

func initStack2() {
	stack2 = make([]int, 0, 20)
}
