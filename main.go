package main

import (
	"encoding/json"
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

// product struct
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

// getProducts handler request get api for products
func getProducts(w http.ResponseWriter, r *http.Request) {
	// cors controll allow
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// if anything else get request comes it will show error
	if r.Method != http.MethodGet {
		http.Error(w, "plz, Give me get request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// create product (post)
func createProduct(w http.ResponseWriter, r *http.Request)  {
	// cors controll allow
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")


	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	// if anything else psot request comes it will show error
	if r.Method != http.MethodPost {
		http.Error(w, "plz, Give me post request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", 400)
		return
	}

	// set values to new product
	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)
	w.WriteHeader(201)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	// Creates a new HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// Registers handler functions
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product",createProduct)

	// Prints a message to the console indicating the server is running
	fmt.Println("==== server running on : 8080 ====")

	// Starts the HTTP server on port 3000 using the mux router
	err := http.ListenAndServe(":8080", mux)

	// If there's an error starting the server, print it
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/0L7DFeQB0UMo_ytgjJ4EPf87gEccQDL6uiyMVzuTI34/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMTgv/OTI3LzE2MS9zbWFs/bC9vcmFuZ2UtZnJ1/aXQtY29sb3ItaWxs/dXN0cmF0aW9uLXBu/Zy5wbmc",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/5_BRcLV8qhdtGjsjGHYkvaWpA0GXvd6V5L7-AL-zxZc/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMjQv/MjE1LzExNC9zbWFs/bC9yaXBlLWFwcGxl/cy1mcnVpdC13aXRo/LWxlYXZlcy1waG90/by5qcGc",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/ZIwSNRtZ5AVAniDBDU3EpGJLMQc_1G-neNobwooYqYE/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9idW5j/aC1iYW5hbmEtZnJ1/aXRzLWN1dC1iYW5h/bmFzLWlzb2xhdGVk/LXdoaXRlLWJhY2tn/cm91bmQtYnVuY2gt/YmFuYW5hLWZydWl0/cy1jdXQtYmFuYW5h/cy1pc29sYXRlZC0x/MDM2NzUzMTkuanBn",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/agNMnjQSLtnTyMIddcz4GiaDYlx2T7G7ORT4EMtu9z4/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMDYv/OTgyLzE1NC9zbWFs/bC9mcmVzaC1yZWQt/b3ItcHVycGxlLWdy/YXBlcy1mcnVpdC1p/c29sYXRlZC1vbi13/aGl0ZS1iYWNrZ3Jv/dW5kLXBob3RvLmpw/Zw",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/PBhQraLJWOy5Aerb6kNwGoRf-uNCFyaUb38CMifp2O0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9tYW5n/by1mcnVpdC1tYW5n/by1jdWJlcy13b29k/ZW4tdGFibGUtNjA1/OTc4ODcuanBn",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Lichi",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/wYwhLuzjJLRovZpxM3gS9mTbale-u7jBx_lN7Hlq2F0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5nZXR0eWltYWdl/cy5jb20vaWQvODAy/ODkzNDI2L3Bob3Rv/L2x5Y2hlZS5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9bVVX/dVZQYkstRENpZEF3/N3JSb2R5dU96a21t/UHpodFlzTHpkSE9L/Q2Ywaz0",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
