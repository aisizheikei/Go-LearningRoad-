package main

import (
	"fmt"
	"time"
)

//程序启动后会创建一个主goroutine去执行
func main() {
	go hello() //go关键字 开启一个单独的goroutine去执行
	fmt.Println("main")
	//main函数执行结束了，最终结果只打出个【main】就结束了

	fmt.Println("-------匿名函数启用goroutine---------------------------------------")
	for i := 0; i <= 1000; i++ {
		go func() {
			fmt.Println(i)
		}()
		//可以看到打出的值并不是0--1000，因为goroutine启用也会耗费时间
	}
	fmt.Println("------匿名函数启用goroutine2----------------------------------------")
	for i := 0; i <= 1000; i++ {
		go func(i int) {
			fmt.Printf("%v、", i)
		}(i)
	}

	time.Sleep(time.Second)
}

func hello() {
	fmt.Println("Hello")
}
