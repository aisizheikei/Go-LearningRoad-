package main

import "fmt"

func main() {
	a := 100
	b := &a
	c := &b
	fmt.Printf("a:%v,b:%v,c:%v\n", a, b, **c)
	fmt.Printf("a:%x,b:%x,c:%x\n", a, b, c)
	fmt.Printf("a:%p,b:%p,c:%p\n", &a, b, c)
}
