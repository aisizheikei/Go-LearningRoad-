package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	getMethod()

	attackRequest("http://www.baidu.com", 300)

}
func getMethod() {
	resp, err := http.Get("http://127.0.0.1:9090/get/test/1")
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(b))
}

func attackRequest(url string, count int) {
	var wg sync.WaitGroup
	temp := 0
	for i := 0; i < count; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("get url failed, err:", err)
				return
			}
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Read resp.Body failed, err:", err)
				return
			}
			fmt.Println(i) //输出结果出乎意料啊。知识点！！！
			temp++
		}()
	}
	wg.Wait()
	fmt.Println("结束了", temp)
}
