// Package controller defines the layer for communication
// between handler and storage
package controller

import (
	"github.com/pkg/errors"
	"github.com/saromanov/grpc-chat/app/service/chat/models"
	"github.com/saromanov/grpc-chat/app/service/chat/storage"
	"github.com/saromanov/grpc-chat/app/service/config"
	"github.com/saromanov/grpc-chat/proto/chat"
)

// Controller defines controller for user and storage
type Controller struct {
	st   *storage.Storage
	conf *config.Config
}

// New creates controller
func New(conf *config.Config) (*Controller, error) {
	st, err := storage.New(conf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to init controller")
	}
	return &Controller{
		st:   st,
		conf: conf,
	}, nil
}

// SendMessage provides inserting of the new message
func (c *Controller) SendMessage(req *chat.Message) error {
	m := &models.Message{
		Body: req.Body,
		User: req.User,
	}

	return c.st.Insert(m)
}

// SearchMessages provides searching of the messages
func (c *Controller) SearchMessages(req *chat.SearchMessagesRequest) ([]*models.Message, error) {
	resp, err := c.st.Search(&models.SearchMessages{
		User: req.User,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to find data")
	}
	return resp, nil
}
