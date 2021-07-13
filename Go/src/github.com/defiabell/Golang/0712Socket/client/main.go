package main

import (
	"fmt"
	"net"
)

//客户端
func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1 failed, err:", err)
		return
	}
	//发送数据
	var msg string
	for {
		fmt.Println("请输入：")
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
