// Package controller defines the layer for communication
// between handler and storage
package controller

import (
	"github.com/pkg/errors"
	"github.com/saromanov/grpc-chat/app/service/chat/storage"
	"github.com/saromanov/grpc-chat/app/service/config"
)


// Controller defines controller for user and storage
type Controller struct {
	st *storage.Storage
	conf *config.Config
}

// New creates controller
func New(conf *config.Config) (*Controller, error) {
	st, err := storage.New(conf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to init controller")
	}
	return &Controller{
		st: st,
		conf: conf, 
	}
}
