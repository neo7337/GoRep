package stub

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) Consume(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received Message : %s", in)
	return &Message{Body: "OK Response From Server!"}, nil
}
