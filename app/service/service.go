// Package service provides initialization of the chat service
package service

import (
	"github.com/pkg/errors"
	"github.com/saromanov/grpc-chat/app/service/chat/handlers"
	"github.com/saromanov/grpc-chat/app/service/chat/storage"
	"github.com/saromanov/grpc-chat/app/service/config"
)

// New provides initialization of the all components of the service
func New(conf *config.Config) error {
	st, err := storage.New(conf)
	if err != nil {
		return errors.Wrap(err, "unable to init storage")
	}
	err = handlers.New()
	if err != nil {
		return errors.Wrap(err, "unable to init handler")
	}
	return nil
}
