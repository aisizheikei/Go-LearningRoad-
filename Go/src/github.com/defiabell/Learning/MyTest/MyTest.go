package main

import "fmt"

func main() {
	testRune()
}

//测试切片使用---证明通过索引赋值给一个变量后是通过值副本的方式，改变变量值不会改变原有的值数据
func testRune() {
	rune1 := []int{1, 2, 3}
	array1 := [...]int{1, 2, 3}
	fmt.Println("rune1--old:", rune1)
	fmt.Println("array1--old:", array1)
	temp := rune1[1]
	temp = 22
	fmt.Println("rune1--new:", rune1)
	temp = array1[1]
	temp = 22
	fmt.Println("array1--new:", array1)
	fmt.Printf("temp:%v\n", temp)
}
