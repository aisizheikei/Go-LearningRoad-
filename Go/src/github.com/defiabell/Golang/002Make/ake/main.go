package main

import "fmt"

func main() {
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a2 := a1
	a1[2] = 13
	fmt.Printf("a1:%v\n", a1)
	fmt.Printf("a2:%v\n", a2)

	m1 := make([]int, 3, 4)
	m1[1] = 12
	m2 := m1
	m2[2] = 13

	fmt.Printf("m1:%v;len(m1):%d;cap(m1):%d\n", m1, len(m1), cap(m1))
	fmt.Printf("m1:%v\n", m2)

	m2 = append(m2, 4, 5)
	fmt.Printf("m2:%v;len(m2):%d;cap(m2):%d\n", m2, len(m2), cap(m2))

	m3 := append(m1, 6)
	fmt.Printf("m3:%v;len(m3):%d;cap(m3):%d\n", m3, len(m3), cap(m3))

	//切片。数组
	a3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	m13 := a3[:]
	m13 = append(m13[:2], m13[3:]...) //删除索引为2的元素
	fmt.Printf("a3:%v\n", a3)         //索引为2的元素删除了，剩下的前移，最后一位不变
	fmt.Printf("m13:%v\n", m13)
}
