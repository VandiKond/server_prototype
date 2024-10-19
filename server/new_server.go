package server

import (
	hash_password "go_server/hash"
	"time"
)

// A server structure
type server struct {
	id                 int
	hash               string
	allowed_time       int
	allowed_operations int
	created_at         time.Time
	// You can add other information to the server
}

// A fuction to create a server
func NewServer(id int, password string, allowed_time int, allowed_operations int) (*server, error) {
	// You can add logic of checking the id or saving the server to the databse

	// Creating the hash of the password
	hash, err := hash_password.HashPassword(password)

	// In case of error returning the error
	if err != nil {
		return nil, err
	}

	// Returning a new server
	return &server{id: id, hash: hash, allowed_time: allowed_time, allowed_operations: allowed_operations, created_at: time.Now()}, nil /* << It hasn't got and error so it's nill */
}
