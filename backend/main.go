package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eyalch/pipeit/backend/code"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	code.NewHandler(r)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
