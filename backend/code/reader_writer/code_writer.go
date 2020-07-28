package readerwriter

import (
	"github.com/eyalch/pipeit/backend/domain"
	"github.com/gorilla/websocket"
)

type codeWriter struct {
	conn *websocket.Conn
}

func NewCodeWriter() domain.CodeWriter {
	return &codeWriter{}
}

func (s *codeWriter) Write(code string) error {
	return s.conn.WriteMessage(websocket.TextMessage, []byte(code))
}
