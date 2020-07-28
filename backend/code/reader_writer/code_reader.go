package readerwriter

import (
	"github.com/eyalch/pipeit/backend/domain"
	"github.com/gorilla/websocket"
)

type codeReader struct {
	conn *websocket.Conn
}

func NewCodeReader() domain.CodeReader {
	return &codeReader{}
}

func (s *codeReader) Read() (string, error) {
	for {
		_, message, err := s.conn.ReadMessage()
		return message, err
	}
}
