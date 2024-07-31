package main

import (
	"fmt"
	"log"

	"meow.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darryn"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	// // request a greeting message
	// message, err := greetings.Hello("Gladys")

	// // we handle the error
	// // if there is an error, we print it to the console and exit the program
	// if err != nil {
	// log.Fatal(err)
	// }

	// // if no error, we print the normal message
	// fmt.Println((message))
}
