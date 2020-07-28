package http

import (
	"log"
	"net/http"

	"github.com/eyalch/pipeit/backend/domain"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type codeHandler struct {
	uc domain.CodeUsecase
}

// NewCodeHandler returns an http.Handler which handles /code routes
func NewCodeHandler(uc domain.CodeUsecase) http.Handler {
	h := codeHandler{uc}

	r := chi.NewRouter()

	r.HandleFunc("/new", h.newCode)
	r.HandleFunc("/pair", h.pair)

	return r
}

var upgrader = websocket.Upgrader{
	// TODO: Remove this
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *codeHandler) newCode(w http.ResponseWriter, r *http.Request) {
	code := h.uc.GetRandomCode()
	log.Println("Generated a new code:", code)

	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	if err = h.uc.SendCodeAndWaitForPair(code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
