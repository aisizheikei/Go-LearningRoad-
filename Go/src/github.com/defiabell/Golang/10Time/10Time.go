package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("时间：", now)

	//时间戳
	fmt.Println("毫秒时间戳：", now.Unix())
	fmt.Println("纳秒时间戳：", now.UnixNano())

	// //定时器
	// timer := time.Tick(time.Second)
	// //每秒执行一次
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	//时间格式化 Go语言不同于其他语言，格式化不是通过'YY-mm-dd'。而是用'2006-01-02 03:04:05'，也就是Go语言的诞生时间
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	//时间计算
	fmt.Println(now.Add(1 * time.Hour))
	timeTemp, err := time.Parse("2006/01/02", "2019/02/20")
	if err != nil {
		fmt.Println("时间转换失败")
	}
	fmt.Println(now.Sub(timeTemp))
	fmt.Println("----------------------------------------------")

	// 根据时区解析时间
	loc, err := time.LoadLocation("")
	if err != nil {
		fmt.Println("load loc failed ,err:", err)
	}
	timeTemp_loc, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/05/22 18:39:00", loc)
	if err != nil {
		fmt.Println("time parase failed. err:", err)
	}
	fmt.Println("时区：", loc, ";时间：", timeTemp_loc)
	fmt.Println("----------------------------------------------")

}
