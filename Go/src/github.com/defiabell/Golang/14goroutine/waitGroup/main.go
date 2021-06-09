package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//wg.Add(10)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	//main函数不知道线程是否都已经结束。所以用wg记录
	wg.Wait()
}

func f1(i int) {
	defer wg.Done()
	fmt.Println(i)
}
