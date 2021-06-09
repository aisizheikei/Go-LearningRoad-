package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("cpu count:%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2) //设置最大goruntine数量

	wg.Add(2)
	go f1()
	go f2()

	wg.Wait()
}

func f1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("1:", i)
	}
}
func f2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("2:", i)
	}
}
