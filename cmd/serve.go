package cmd

import (
	globalroute "ecommerce/global_route"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"      // Importing the fmt package for formatted I/O
	"net/http" // Importing the net/http package to build HTTP servers
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

// // cors handler
// func handleCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Hasib")
// }

// // handle options(preflight request)
// func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(200)
// 	}
// }

// // middleware (currently we are not using middleware as we are doing it globally.)
// func corsHandler(next http.Handler) http.Handler {
// 	handleCors := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Hasib")
// 		next.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(handleCors)
// }

func Serve() {
	// Creates a new HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// Registers handler functions

	//  ===handle before go 1.20===
	// mux.HandleFunc("/", helloHandler)
	// mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/products", getProducts) // this is example of how we used to declare routes before go 1.20

	// ===advance handler without middleware (we need to handle OPTIONS(preflight) manualy)===
	// mux.Handle("GET /products", http.HandlerFunc(getProducts))
	// mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))

	// mux.Handle("POST /create-product", http.HandlerFunc(createProduct))
	// mux.Handle("OPTIONS /create-product", http.HandlerFunc(createProduct))

	// ===advance handler with middleware===

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productID}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	// global route
	globalRouter := globalroute.GlobalRouter(mux)

	// Prints a message to the console indicating the server is running
	fmt.Println("==== server running on : 8080 ====")

	// Starts the HTTP server on port 3000 using the mux router
	err := http.ListenAndServe(":8080", globalRouter)

	// If there's an error starting the server, print it
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
