package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eyalch/pipeit/backend/code"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Handlers
	r.Mount("/code", code.NewHandler())

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
