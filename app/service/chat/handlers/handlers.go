// Package handlers defines handling of messages via gRPC
package handlers

package delivery

import (
	"github.com/saromanov/grpc-chat"
	"github.com/saromanov/grpc-chat/proto/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ChatServer struct {
	ChatController chat.Controller
}

// New provides initialization of the gRPC handler
func New(serverInstance *grpc.Server, chatController chat.Controller) {
	server := &ChatServer{
		ChatController: chatController,
	}

	chat.RegisterChatServer(serverInstance, server)
}

func (s *ChatServer) AddMessage(req *chat.Message, res *chat.MessageResponse) error {
	username := getUsernameFromContext(ctx)
	server.ChatController.SendMessage(username, req)
	return new(chatPB.SendMessageResponse), nil
}

func (s *ChatServer) SerchMessages(ctx context.Context, req *chat.ClientMessage) (*chatPB.SendMessageResponse, error) {
}