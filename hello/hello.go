package main

import (
	"fmt"
	"log"

	"github.com/AraiYuno/go-tutorial/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message for the names.
	messages, err := greetings.Hellos(names)
	// if an error was returned, print it to the console and
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
}
