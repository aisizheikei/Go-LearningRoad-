package main

import "fmt"

func main() {
	i1 := 10
	//var p1 *int //nil pointer
	var p1 *int = &i1 //nil pointer
	*p1 = 100
	fmt.Println(&i1, i1)
	fmt.Println(p1)

}
