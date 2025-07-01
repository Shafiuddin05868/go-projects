package main

import (
	"fmt"
	"io"
	web "net/http"
)

func main() {
	fmt.Println("Hellow, Go standard library")
	res, err := web.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	} else {
		fmt.Println("Status ", res.Status)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(body))
}
