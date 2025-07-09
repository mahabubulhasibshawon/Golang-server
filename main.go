package main

import (
    "fmt"       // Importing the fmt package for formatted I/O
    "net/http"  // Importing the net/http package to build HTTP servers
)

// helloHandler handles requests to the root URL ("/")
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
}

// aboutHandler handles requests to the "/about" URL
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    // Writes information about the developer to the HTTP response
    fmt.Fprintln(w, "I'm Mahabub Ul Hasib Shawon\nI'm a software engineer\nI love golang for it's goroutine")
}

func main() {
    // Creates a new HTTP request multiplexer (router)
    mux := http.NewServeMux()

    // Registers handler functions
    mux.HandleFunc("/", helloHandler)
    mux.HandleFunc("/about", aboutHandler)

    // Prints a message to the console indicating the server is running
    fmt.Println("==== server running on : 3000 ====")

    // Starts the HTTP server on port 3000 using the mux router
    err := http.ListenAndServe(":3000", mux)

    // If there's an error starting the server, print it
    if err != nil {
        fmt.Println("Error starting the server", err)
    }
}
