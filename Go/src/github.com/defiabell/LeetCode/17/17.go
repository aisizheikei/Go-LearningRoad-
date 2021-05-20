package main

import (
	"fmt"
)

func main() {
	var digits string
	digits = "23"
	fmt.Println("结果是：", letterCombinations(digits))
}

var digitsMap = map[string][]string{
	"1": {""},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {

	var result = make([]string, 0, 100)
	for i, v := range digits {
		if i == 0 {
			for _, v1 := range digitsMap[string(v)] {
				result = append(result, v1)
			}
		} else {
			result = getResult(result, string(v))
		}

	}
	return result
}

func getResult(str []string, key string) (resultList []string) {

	resultList = make([]string, 0, 100)
	for _, v := range digitsMap[key] {
		for _, v2 := range str {
			resultList = append(resultList, v2+v)
		}
	}
	return
}
