package main

import "fmt"

func main() {
	a := 100
	b := &a
	c := &b
	fmt.Printf("a:%v,b:%v,c:%v\n", a, b, **c)
	fmt.Printf("a:%x,b:%x,c:%x\n", a, b, c)
	fmt.Printf("a:%p,b:%p,c:%p\n", &a, b, c)
	fmt.Println("----------------------------------------------")
	pointVal()
}

func pointVal() {
	m := make(map[int]int, 3)
	m[1] = 1
	m[2] = 2

	fmt.Printf("m的值:%v;m的地址：%p", m, m)
	fmt.Println()
	fmt.Printf("&m的值%v,地址：%p\n", &m, &m)
	fmt.Println()
	fmt.Printf("m的len:%v;容量\n", len(m))
}
