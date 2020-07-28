package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eyalch/pipeit/backend/code"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gomodule/redigo/redis"
	"github.com/streadway/amqp"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func main() {
	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer amqpConn.Close()

	amqpCh, err := amqpConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer amqpCh.Close()

	redisConn, err := redis.DialURL("redis://localhost:6379")
	failOnError(err, "Failed to connect to Redis")
	defer redisConn.Close()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Handlers
	r.Mount("/code", code.NewHandler(amqpCh))

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
