package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thashinsharon/Go_with_Habib/Lec51/database"
	"github.com/thashinsharon/Go_with_Habib/Lec51/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valied json", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	util.SendData(w, newProduct, 200)
}
