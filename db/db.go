package db

import (
	"log"
	"os"

	"github.com/dbongo/app/db/storage"
	"gopkg.in/mgo.v2"
)

const (
	// DefaultDBURL ...
	DefaultDBURL = "127.0.0.1:27017"
	// DefaultDBName ...
	DefaultDBName = "appdb"
)

// Storage ...
type Storage struct {
	*storage.Storage
}

func conn() (*storage.Storage, error) {
	url := os.Getenv("MONGO_ADDRESS")
	if url == "" {
		url = DefaultDBURL
	}
	name := os.Getenv("MONGO_DATABASE")
	if name == "" {
		name = DefaultDBName
	}
	log.Printf("db connected to %s at %s", name, url)
	return storage.Open(url, name)
}

// Conn ...
func Conn() (*Storage, error) {
	var (
		strg Storage
		err  error
	)
	strg.Storage, err = conn()
	return &strg, err
}

// Users returns the users collection from MongoDB.
func (s *Storage) Users() *storage.Collection {
	emailIndex := mgo.Index{Key: []string{"email"}, Unique: true}
	c := s.Collection("users")
	c.EnsureIndex(emailIndex)
	return c
}
