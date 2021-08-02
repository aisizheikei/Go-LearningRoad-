package main

import (
	"fmt"
)

func main() {
	str := "str"
	fmt.Println("test_string:", test_string(str))
	str = "str"
	fmt.Println("test_string_point:", test_string_point(&str))
}
func test_string_point(str *string) string {
	if len(*str) <= 2 {
		return *str
	} else {
		*str = "sad"
	}
	for i, c := range *str {
		fmt.Println("i:", i, ";c:", string(c))
	}
	return *str
}
func test_string(str string) string {
	if len(str) <= 2 {
		return str
	} else {
		str = "sad"
		return str
	}
}
