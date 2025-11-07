package handlers

import (
	"net/http"

	"Lec53/database"

	"Lec53/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 201)
}
