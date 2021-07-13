package main

import "fmt"

func main() {
	var str1 string
	fmt.Scan(&str1)
	fmt.Println("str1:", str1)

	var (
		name string
		age  int
	)

	_, err := fmt.Scanf("%s %d\n", &name, &age)
	fmt.Printf("name:%s,age:%d,err:%v\n", name, age, err)

	var str3 string
	fmt.Scanln(&str3)
	fmt.Println("str3:", str3)

}
