package main

import "fmt"

var (
	name string
	age  int
)

const (
	a1 = iota
	a2
	a3 = 1
	a4 = iota
)

func main() {
	fmt.Println("Hello world")
	fmt.Printf("你好%s%s", "中国", "美国")
	fmt.Println()
	name = "jk&jy"
	age = 1314
	fmt.Printf("My name is %s, I`m %c years old", name, age)
	fmt.Println()
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(a1, a2, a3, a4)
	str1 := `
	窗前明月光，
	疑是地上霜。
	举头望明月，
	低头思故乡。
	thanks
	`
	fmt.Println(str1, "str1长度", len(str1))
	str1_arr := []rune(str1)
	fmt.Println((str1_arr))
	fmt.Println((str1_arr[27:]))

	str2 := "你好汉字"
	fmt.Printf("%x:%v\n", &str2, str2)
	str2 = "ss你好汉字"
	fmt.Printf("%x:%v\n", &str2, str2)
	str3 := str2
	fmt.Printf("%x:%v\n", &str3, str3)
	str3 = "我是啥"
	fmt.Printf("%x:%v\n", &str2, str2)

	fmt.Println("str2长度：", len(str2))

}
