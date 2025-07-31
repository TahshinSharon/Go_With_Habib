package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Tahshin. I'm Youtuber. I'm Software Engineer")
}
func main() {
	router := http.NewServeMux()

	router.HandleFunc("/hello", helloHandler)

	router.HandleFunc("/about", aboutHandler)

	fmt.Println("Server Running On :3000")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Panicln("Error Starting The Server", err)
	}
}
