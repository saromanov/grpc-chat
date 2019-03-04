// Package storage provides implementation of Postgesql storage
package storage

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saromanov/grpc-chat/app/service/chat/models"
)

// Storage implements db handling with Postgesql
type Storage struct {
	db *gorm.DB
}

// Create provides init for postgesql storage
func Create(s *structs.Config) (st.Storage, error) {
	if s == nil {
		return nil, errors.New("config is not defined")
	}
	args := "dbname=selitra"
	if s.DBName != "" && s.DBPassword != "" && s.DBUser != "" {
		args += fmt.Sprintf(" user=%s dbname=%s password=%s", s.DBUser, s.DBName, s.DBPassword)
	}
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}
	db.AutoMigrate(&st.LogRequest{})
	return &storage{
		db: db,
	}, nil
}

// Insert provides inserting of data
func (s *storage) Insert(m *models.Message) error {
	err := s.db.Create(m).Error
	if err != nil {
		return fmt.Errorf("storage: unable to insert data: %v", err)
	}
	return nil
}

// Insert provides finding data
func (s *storage) Search(sr *models.SearchRequest) ([]*models.Message, error) {
	var response []*models.Message
	err := s.makeQuery(s.db, sr).Find(&response).Error
	if err != nil {
		return nil, fmt.Errorf("storage: unable to find data: %v", err)
	}
	return response, nil
}

// makeQuery provides making of the query to Postgresql
func (s *storage) makeQuery(db *gorm.DB, sr *modles.SearchRequest) *gorm.DB {
	if sr.User != "" {
		db = db.Where("user=?", sr.User)
		return db
	}
	return db
}

// Close provides closing of db
func (s *storage) Close() error {
	return s.db.Close()
}