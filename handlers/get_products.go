package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

// getProducts handler request get api for products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// cors controll allow
	// handleCors(w)
	// if r.Method == http.MethodOptions {
	// 	handlePreflightReq(w, r)
	// 	return
	// }

	// if anything else get request comes it will show error
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "plz, Give me get request", 400)
	// 	return
	// }
	util.SendData(w, database.ProductList,200)
}
