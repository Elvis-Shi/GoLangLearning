package main

import  (
    "fmt"
    "log"
    "ssdd.com/greating"
)

import "rsc.io/quote"

func main() {
    log.SetPrefix("greetings: ");
    log.SetFlags(0);
    
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    message, err := greetings.Hello("Elvis")
    if err != nil {
        log.Fatal(err);       
    }

    fmt.Println(message)


    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}