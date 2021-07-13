package main

import (
	"fmt"
	"sync"
)

//锁
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(2)
	go addx()
	go addx()
	wg.Wait()
	fmt.Println("end x:", x)
}

func addx() {
	defer wg.Done()
	lock.Lock()
	for i := 0; i < 500000; i++ {
		x += 1
	}
	lock.Unlock()
	fmt.Println("addx:", x)
}
