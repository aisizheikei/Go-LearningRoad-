package main

import (
	"fmt"

	. "github.com/defiabell/NiuKe/0621LRU_Struct/LRU_func"
	. "github.com/defiabell/NiuKe/0621LRU_Struct/LRU_func2"
)

func main() {
	args1 := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	result1 := LRU(args1, 3)
	result2 := LRU2(args1, 3)
	fmt.Println("result1:", result1) //[1,-1]
	fmt.Println("result1:", result2) //[1,-1]

}
