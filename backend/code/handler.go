package code

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var upgrader = websocket.Upgrader{
	// TODO: Remove this
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handler handles /code routes
type handler struct {
	redisConn   *redis.Conn
	amqpChannel *amqp.Channel
}

// NewHandler returns an http.Handler that handles /code routes
func NewHandler(redisConn *redis.Conn, amqpChannel *amqp.Channel) http.Handler {
	h := handler{redisConn, amqpChannel}
	return h.router()
}

func (h *handler) router() http.Handler {
	r := chi.NewRouter()

	r.HandleFunc("/new", h.newCode)
	r.HandleFunc("/pair", h.pair)

	return r
}

func (h *handler) newCode(w http.ResponseWriter, r *http.Request) {
	// Create a new code
	code := h.store.create()
	log.Println("Generated a new code:", code)

	q, err := h.amqpChannel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Send the code
	err = conn.WriteMessage(websocket.TextMessage, []byte(code))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// for {
	// 	// Read a message
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Print("read:", err)
	// 		break
	// 	}

	// 	// Print the received message
	// 	log.Printf("recv: %s", message)
	// }
}

func (h *handler) pair(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	// Pop the code from the store
	if err := h.store.pop(code); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	q, err := h.amqpChannel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
}
