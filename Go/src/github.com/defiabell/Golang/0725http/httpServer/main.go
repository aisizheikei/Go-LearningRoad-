package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/posts/Go/defiabell/", f1)
	http.HandleFunc("/get/test/", f2)

	http.ListenAndServe("127.0.0.1:9090", nil) //http://127.0.0.1:9090/posts/Go/defiabell/

}

func f1(w http.ResponseWriter, r *http.Request) {
	str := "<h1>Hello Defiabell!</h1>"
	w.Write([]byte(str))
}
func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Test!!!"))
}
