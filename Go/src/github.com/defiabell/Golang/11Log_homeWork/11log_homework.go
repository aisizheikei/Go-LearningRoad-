package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("这是一个日志1")
	testLog()
	log.Println("这是一个日志2")

}

func testLog() {
	fileObj, err := os.OpenFile("./11log_homwork.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed. err:", err)
	}
	//指定日志输出位置
	log.SetOutput(fileObj)
}
