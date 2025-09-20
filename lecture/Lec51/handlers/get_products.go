package handlers

import (
	"net/http"

	"github.com/thashinsharon/Go_with_Habib/Lec51/database"
	"github.com/thashinsharon/Go_with_Habib/Lec51/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 201)
}
