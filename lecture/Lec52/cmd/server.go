package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thashinsharon/Go_with_Habib/Lec51/middleware"
	"github.com/thashinsharon/Go_with_Habib/Lec51/util"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)
	router := http.NewServeMux()
	initRoutes(router, manager)
	globalRouter := util.GlobalRouter(router)
	fmt.Println("Server Running On :3000")
	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		log.Panicln("Error Starting The Server", err)
	}
}
