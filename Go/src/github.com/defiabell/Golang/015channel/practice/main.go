package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var c1 chan int
	c1 = make(chan int, 5)
	wg.Add(1)
	go channelfunc1(c1)
	for {
		x, ok := <-c1
		if !ok {
			fmt.Println("read end")
			break
		}
		fmt.Println(x)
	}
	wg.Wait()

}

func channelfunc1(c1 chan<- int) {
	defer wg.Done()
	fmt.Println("channel action:")
	for i := 1; i < 10; i++ {
		c1 <- i
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Printf("channel End")
	close(c1) //不关闭的话，通道没数据 的时候再读取就会堵塞
}
