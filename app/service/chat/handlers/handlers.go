// Package handlers defines handling of messages via gRPC
package handlers

package delivery

import (
	"errors"
	"github.com/saromanov/grpc-chat"
	"github.com/saromanov/grpc-chat/proto/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ChatServer defines struct for chat server handling
type ChatServer struct {
	ChatController chat.Controller
}

// New provides initialization of the gRPC handler
func New(server *grpc.Server, chatController chat.Controller) error {
	if server == nil {
		return errors.New("server is not defined")
	}
	s:= &ChatServer{
		ChatController: chatController,
	}

	chat.RegisterChatServer(s, server)
	return nil
}

func (s *ChatServer) AddMessage(req *chat.Message, res *chat.MessageResponse) error {
	username := getUsernameFromContext(ctx)
	server.ChatController.SendMessage(username, req)
	return new(chatPB.SendMessageResponse), nil
}

func (s *ChatServer) SerchMessages(ctx context.Context, req *chat.ClientMessage) (*chatPB.SendMessageResponse, error) {
}