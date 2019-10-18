package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.ListenAndServe(":8080", nil) // set listen port
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorldHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello New World2!!")
}
