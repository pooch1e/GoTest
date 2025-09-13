package main

import (
	"fmt"

	"example/greetings"
	"log"
)

func main() {
	message, err := greetings.Hello(" ")
	if err != nil {
		return log.Fatal(err)
	}
	fmt.Println(message)

}
