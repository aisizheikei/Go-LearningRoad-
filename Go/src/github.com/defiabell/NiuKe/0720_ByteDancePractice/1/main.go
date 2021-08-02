package main

import "fmt"

func main() {
	var count int
	fmt.Scanln(&count)
	result := make([]string, 0)
	for i := 0; i < count; i++ {
		var str string
		fmt.Scanln(&str)
		result = append(result, reBuild(str))
	}
	for _, val := range result {
		fmt.Println(val)
	}
}
func reBuild(str string) string {
	if len(str) <= 2 {
		return str
	}
	var res string = str[:2]
	for i, val := range str {
		length := len(res)
		if i < 2 {
			continue
		} else if str[i] == res[length-1] {
			length := len(res)
			if str[i] == res[length-2] {
				continue
			} else if length > 3 && res[length-2] == res[length-3] {
				continue
			}
		}
		res += string(val)
	}
	return res
}
