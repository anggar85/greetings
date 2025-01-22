package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name esta vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Helos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {

		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"Hola %v, como estas?",
		"Que onda %v, como estas?",
		"Heyts %v, has regresado?",
	}
	return formats[rand.Intn(len(formats))]
}
