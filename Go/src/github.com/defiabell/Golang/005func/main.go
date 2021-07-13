package main

import "fmt"

//go语言没有默认参数的概念
func main() {

	fmt.Println("sum4", sum4(1, 2, 3, 4, 5))
}

//没有返回值的
func sum1(x int, y int) {
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

//声明返回值的
func sum2(x, y int) (res int) {
	res = x + y
	return //可以不指定res
}

//未声明返回值的
func sum3(x, y int) int {
	return x + y
}

//可变参数
func sum4(x ...int) int {
	var res int
	for _, v := range x {
		res += v
		//fmt.Printf("%d:%d\n", i, res)
	}
	return res

}
