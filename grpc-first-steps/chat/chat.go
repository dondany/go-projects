package chat

import (
	context "context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body: %s", message.Body)
	return &Message{Body: "Yo from the server"}, nil
}