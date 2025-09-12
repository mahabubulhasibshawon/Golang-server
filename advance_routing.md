# ===handle before go 1.20===
```
 mux.HandleFunc("/", helloHandler)
mux.HandleFunc("/about", aboutHandler)
mux.HandleFunc("/products", getProducts) // this is example of how we used to declare routes before go 1.20
```

# ===advance handler without middleware (we need to handle OPTIONS(preflight) manualy)===
```	

	 mux.Handle("GET /products", http.HandlerFunc(getProducts))
	 mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))
```

```
	// mux.Handle("POST /create-product", http.HandlerFunc(createProduct))
	// mux.Handle("OPTIONS /create-product", http.HandlerFunc(createProduct))

	// ===advance handler with middleware===
	mux.Handle("GET /products", corsHandler(http.HandlerFunc(getProducts)))
	mux.Handle("POST /create-product", corsHandler(http.HandlerFunc(createProduct)))
```

---

### Router Setup

```go
mux := http.NewServeMux()
mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
mux.Handle("POST /createTask", http.HandlerFunc(createTask))
mux.Handle("GET /getTask", http.HandlerFunc(GetTask))
mux.Handle("OPTIONS /getTask", http.HandlerFunc(GetTask))
```

---

### CORS Handler

```go
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
	w.Header().Set("Content-Type", "application/json")
}
```

---

### Preflight Request Handler

```go
func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}
```

---

### Get Task Handler

```go
func GetTask(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusOK) // handle preflight without extra function
		return
	}
	sendData(w, productList, http.StatusOK)
}
```

---

### Create Task Handler

```go
func createTask(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	sendData(w, productList, http.StatusOK)
}
```

---

### Middleware Example

```go
func CorsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleCors)
	return handler
}
```

---

### Global Router with Preflight Handling

```go
func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
```

---

### Server Entry Point

```go
func main() {
	mux := http.NewServeMux()

	fmt.Println("Server running on: 8080")
	globalRouter := globalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
```

---

