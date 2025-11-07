package middleware

import "net/http"

func GlobalRouter(router *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		// Hanle cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Habib")
		w.Header().Set("Content-Type", "application/json")

		//Handle Preflight Request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		router.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
