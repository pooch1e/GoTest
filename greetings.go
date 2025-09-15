package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func randomFormat() string {
	formats := []string{
		"Hi %v, nice to see you",
		"welcome %v",
		"Test 3, %v",
	}
	println(rand.Intn(len(formats)) + 1)
	return formats[rand.Intn(len(formats))]
}
