package main

import (
	"fmt"
	"log"

	"alex.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Alessandro")
	fmt.Println(message)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// A slice of names.
	names := []string{"Alessandro", "Luigi", "Gaetano"}
	message2, err := greetings.Hellos(names)
	fmt.Println(message2)

	//Get the error
	message, err = greetings.Hello("")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
