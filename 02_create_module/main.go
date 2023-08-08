package main

import (
	"fmt"
	"go-dev-tutorial/02_create_module/hello"
)

func main() {
	message, err := hello.SayHello("John Doe")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}
