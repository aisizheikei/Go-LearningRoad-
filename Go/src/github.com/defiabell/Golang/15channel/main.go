package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var b chan int

func main() {

	noCacheChannelfunc()

	wg.Wait()
	close(b)
}
func noCacheChannelfunc() {
	b = make(chan int) //不带缓冲区的channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Printf("无缓冲区channel后台从通道中获取到了值：%v\n", x)
	}()
	b <- 10

}
