// package models contains models for chat
package models

import "github.com/saromanov/grpc-chat/proto/chat"

// Message defines struct for messaging
type Message struct {
	ID string
	Body string
	User string
}

// ToProto provides converting of the message to
// protobuf representation
func (m *Message) ToProto()*chat.Message {
	return &chat.Message {
		Body: m.Body,
		User: m.User,
	}
}