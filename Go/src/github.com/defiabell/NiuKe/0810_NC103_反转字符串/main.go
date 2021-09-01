/*
描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）
示例1
输入："abcd"
返回值："dcba"
*/
package main

import "fmt"

func main() {
	args1 := "asdjlkafjrnv"
	fmt.Println(solve(args1))
}

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	// write code here
	len := len(str)
	result := make([]rune, len)
	for i, v := range str {
		result[len-i-1] = v
	}
	return string(result)
}
