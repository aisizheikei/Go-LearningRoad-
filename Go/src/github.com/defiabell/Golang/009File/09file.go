/*
文件中插入 数据
1.打开文件
2.读取插入前的数据 写入 临时文件
3.插入数据
4.循环剩下的数据写入 临时文件
5.临时文件 覆盖 文件

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//Open '09txt.txt'
	fileObj, err := os.Open("./09txt.txt")
	if err != nil {
		fmt.Println("open file failed!")
	}
	//close the file
	defer fileObj.Close()
	//创建一个临时文件，
	tempFile, err := os.OpenFile("./09file.tmp", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("create file failed")
	}
	defer tempFile.Close()

	reader := bufio.NewReader(fileObj)
	tmpLine, err := reader.ReadString('\n')
	if err == io.EOF {
		fmt.Println("the file is end")
	}
	if err != nil {
		fmt.Println("read from file failed, err:", err)
	}
	tempFile.Write([]byte(tmpLine))

	x := []byte{'a'}
	//tempFile.Seek(8, 0)//移动光标到指定位置
	tempFile.Write(x) //插入的数据
	//将剩下的数据写入临时文件
	for {
		var x [1024]byte
		n, err := reader.Read(x[:])
		if err == io.EOF {
			tempFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("read file failed")
			return
		}
		tempFile.Write(x[:n])
	}
	tempFile.Close()
	fileObj.Close()
	//os.Rename("./09file.tmp", "./09file.txt")
	fmt.Println("over")
}
