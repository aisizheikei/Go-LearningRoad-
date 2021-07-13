package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"Joke","age":15}`
	var p person
	json.Unmarshal([]byte(str), &p) //结构体是值类型
	fmt.Println(p)
}
