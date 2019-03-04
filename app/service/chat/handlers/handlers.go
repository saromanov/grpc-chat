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

	chatPB.RegisterChatServer(serverInstance, server)
}

func (server *ChatServer) AddMessage(req *chat.Message, res *chat.MessageResponse) error {
	return nil
}

func (server *ChatServer) SendMessage(ctx context.Context, req *chatPB.ClientMessage) (*chatPB.SendMessageResponse, error) {
	return nil
}