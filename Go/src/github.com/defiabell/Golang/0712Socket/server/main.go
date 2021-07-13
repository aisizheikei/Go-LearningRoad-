package main

import (
	"fmt"
	"net"
)

//服务端
func main() {
	//1、启动本地端口服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:", err)
		return
	}
	for {
		//2、等待别人来建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		//3、与客户端通信
		go processConn(conn)
	}
}
func processConn(conn net.Conn) {
	//3、与客户端通信
	var tmp [128]byte
	for { //能重复接收一个客户端多次请求
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
