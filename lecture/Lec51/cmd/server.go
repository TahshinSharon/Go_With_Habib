package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thashinsharon/Go_with_Habib/Lec51/handlers"
	"github.com/thashinsharon/Go_with_Habib/Lec51/util"
)

func Serve() {
	router := http.NewServeMux()

	router.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	router.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))
	router.Handle("GET /products/{Id}", http.HandlerFunc(handlers.GetProductById))

	globalRouter := util.GlobalRouter(router)
	fmt.Println("Server Running On :3000")
	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		log.Panicln("Error Starting The Server", err)
	}
}
