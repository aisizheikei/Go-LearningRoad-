//读写锁 与 锁
package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0
var lock sync.Mutex
var rwLock sync.RWMutex
var wg sync.WaitGroup

func main() {
	sync_Mutex()
	sync_RWMutex()
}

func sync_Mutex() {
	x = 0
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go write_lock()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read_lock()
	}
	wg.Wait()
	fmt.Println("sync_mutex用时：", (time.Now().Sub(start)), ";x:", x)
}
func sync_RWMutex() {
	x = 0
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go write_rwLock()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read_rwLock()
	}
	wg.Wait()
	fmt.Println("sync_mutex用时：", (time.Now().Sub(start)), ";x:", x)
}
func write_lock() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}
func read_lock() {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}

func write_rwLock() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}
func read_rwLock() {
	defer wg.Done()
	rwLock.RLock()
	time.Sleep(time.Millisecond * 2)
	rwLock.RUnlock()
}
