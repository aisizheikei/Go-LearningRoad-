package main

import "fmt"

func main() {
	var m1 map[string]int //nil
	fmt.Println(m1, m1 == nil)

	m1 = make(map[string]int, 6)
	fmt.Println(m1, m1 == nil)
	m1["小明"] = 10
	m1["小红"] = 12
	v, ok := m1["小张"]
	if ok {
		fmt.Printf("%v\n", v)
	} else {
		fmt.Printf("查无此人\n")
	}
	//map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除
	delete(m1, "小明")
	fmt.Println(m1)
}
