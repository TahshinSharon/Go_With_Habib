package handlers

import (
	"net/http"
	"strconv"

	"github.com/thashinsharon/Go_with_Habib/Lec51/database"
	"github.com/thashinsharon/Go_with_Habib/Lec51/util"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Data not found", 404)
}
