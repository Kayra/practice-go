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
		"Hi %v, welcome!",
		"Great to see you again %v!",
		"How are you today %v?",
	}
	return formats[rand.Intn(len(formats))]
}
