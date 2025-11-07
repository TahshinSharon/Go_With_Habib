package cmd

import (
	"fmt"
	"log"
	"net/http"

	"Lec53/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	router := http.NewServeMux()
	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Hudai,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(router)
	initRoutes(router, manager)
	fmt.Println("Server Running On :3000")
	err := http.ListenAndServe(":3000", wrappedMux)

	if err != nil {
		log.Panicln("Error Starting The Server", err)
	}
}
