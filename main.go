package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck)

	fmt.Println("listening on 8000")
	http.ListenAndServe(":8000", router)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("aok")
}
