package main

import (
	"fmt"
	"net"
	"strconv"
)

//客户端
func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1 failed, err:", err)
		return
	}
	//连续发送数据
	sendMessageRepeat(conn)
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

func sendMessageRepeat(conn net.Conn) {
	for i := 0; i < 20; i++ {
		conn.Write([]byte("发送:" + strconv.Itoa(i)))
		//time.Sleep(time.Second) //sleep1秒就不会有tcp黏包现象了
	}
}
