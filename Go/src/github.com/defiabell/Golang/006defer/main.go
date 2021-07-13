package main

import "fmt"

func main() {

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //返回的是5
}
func f2() (x int) {
	defer func() {
		x++ //修改的是x，是返回值
	}()
	return 5 //返回的是6。第一步赋值x=5，第二步++，第三部返回
}
func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //返回的是y=x=5。
}
func f4() (x int) {
	defer func(x int) {
		x++ //参数x是副本。
	}(x)
	return 5 //返回的值是5。
}
