package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request)  {
	productID := r.PathValue("productID")

	pId , err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid ID",400)
		return
	}
	
	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "data not found", 404)
}