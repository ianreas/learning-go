package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// a function that starts with a capital letter can be exported and called from somewhere else
// go functions can return multiple values
func Hello(name string) (string, error) { // returns two values (not string OR error but string AND error)
	// If no name was given, return an error with a message.
	if name == "" {
		// return an empty message and an error
		return "", errors.New("empty name")
	}

	// return  a greeting that embeds the name in a message
	// := is a shorthand for declaring and initializing a variable
	// long version:
	// var message string
	// is var a const or a let?
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(RandomFormat(), name) // Sprintf substitutes %v with name

	// no error, so for the second value (error), return nil
	return message, nil
}


// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	// in go, you create maps using the following syntax: 
	//  make(map[key-type]value-type)
	messages := make(map[string]string)
	// loop through the received slice of names, calling
	// the hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// in the map, associate the retriveed message with the name
		messages[name] = message

	}
	return messages, nil

}

func RandomFormat() string {
	// curly bracket array syntax?
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// i am guessting rand.Intn(len(formats)) returns a random number from 0 to length of the formats array
	// and then formats[number] just returns the element at that index
	log.Println(rand.Intn(len(formats)))
	return formats[rand.Intn(len(formats))]
}
