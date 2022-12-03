// A unit test is a function that tests a specific piece of code from a program or package.
// The job of unit tests is to check the correctness of an application, and they are a crucial part of the software

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/endyApina/golang-meetup/greetings"
	"github.com/endyApina/golang-meetup/http/client"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Adesina", "David", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	//router
	http.HandleFunc("/hello", client.GetEmployeeHandler)
	http.ListenAndServe(":8090", nil)
}
