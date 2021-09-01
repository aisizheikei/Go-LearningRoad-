package main

import (
	"flag"
	"fmt"
)

//0805flag.exe -age 20 -name 小李
func main() {
	age := flag.Int("age", 100, "请输入年龄") //age参数，默认值100，后面是描述；返回的是指针
	fmt.Println(*age)                    //解析之前获取不到用户输入的值
	var name string
	flag.StringVar(&name, "name", "张三", "请输入名字") //name参数，默认值张三，后面是描述
	flag.Parse()                                 //解析命令传入参数，赋值给对应变量
	fmt.Println(*age, name)
}
