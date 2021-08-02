package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err：", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入：")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入失败，err:", err)
			return
		}
		socket.Write([]byte(msg))
		//接收回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("Read reply message failed，err：", err)
			return
		}
		fmt.Println("收到的回复信息：", string(reply[:n]))
	}
}
