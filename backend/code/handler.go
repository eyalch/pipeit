package code

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handler handles /code routes
type handler struct {
	store store
}

// NewHandler returns an http.Handler that handles /code routes
func NewHandler(r *mux.Router) {
	h := handler{store{}}

	s := r.PathPrefix("/code").Subrouter()
	s.HandleFunc("/new", h.newCodeHandler)
}

func (h *handler) newCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	// Create a new code
	code := h.store.create()
	log.Println("Generated a new code:", code)

	// Send the code
	err = conn.WriteMessage(websocket.TextMessage, []byte(code))
	if err != nil {
		log.Print("write:", err)
		return
	}
}
