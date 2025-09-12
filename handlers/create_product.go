package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

// create product (post)
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// cors controll allow
	// handleCors(w)

	// == here we were handling ooptions but now we are handling options from global route. that means whatever the reqest is it'll handle options first then it will move forward. so options or preflight hopefully wouldn't bother again
	// if r.Method == http.MethodOptions {
	// 	handlePreflightReq(w, r)
	// 	return
	// }

	// if anything else psot request comes it will show error
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "plz, Give me post request", 400)
	// 	return
	// }

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", 400)
		return
	}

	// set values to new product
	newProduct.ID = len(database.ProductList) + 1

	database.ProductList=append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
