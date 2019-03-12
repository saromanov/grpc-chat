// Package handlers defines handling of messages via gRPC
package handlers

import (
	"errors"

	"github.com/saromanov/grpc-chat/app/service/chat/controller"
	"github.com/saromanov/grpc-chat/proto/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ChatServer defines struct for chat server handling
type ChatServer struct {
	cont controller.Controller
}

// New provides initialization of the gRPC handler
func New(server *grpc.Server, chatController controller.Controller) error {
	if server == nil {
		return errors.New("server is not defined")
	}
	s := &ChatServer{
		cont: chatController,
	}

	chat.RegisterChatServer(server, s)
	return nil
}

// AddMessage provides adding of the new message to db
func (s *ChatServer) AddMessage(ctx context.Context, req *chat.Message) (*chat.MessageResponse, error) {
	return nil, nil
}

// SearchMessages provides searching of the messages
func (s *ChatServer) SearchMessages(ctx context.Context, req *chat.SearchMessagesRequest) (*chat.MessagesResponse, error) {
	return nil, nil
}
