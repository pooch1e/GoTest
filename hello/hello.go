package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message, err := greetings.Hello("Joel")
	if err != nil {
		println("error, no name provided")
	}
	fmt.Println(message)

}
