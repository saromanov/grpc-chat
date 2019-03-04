// Package service provides initialization of the chat service
package service

import (
	"net"

	"github.com/pkg/errors"
	"github.com/saromanov/grpc-chat/app/service/chat/handlers"
	"github.com/saromanov/grpc-chat/app/service/chat/storage"
	"github.com/saromanov/grpc-chat/app/service/config"
	"google.golang.org/grpc"
)

// New provides initialization of the all components of the service
func New(conf *config.Config, lis net.Listener) (chan bool, error) {
	st, err := storage.New(conf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to init storage")
	}

	grpcServer := grpc.NewServer()
	err = handlers.New(grpcServer, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to init handler")
	}
	done := make(chan bool)
	go func(l net.Listener) {
		grpcServer.Serve(lis)
		done <- true
	}(lis)
	return done, nil
}
