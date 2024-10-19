package server

import (
	"fmt"
	hash_password "go_server/hash"
	"time"
)

// Structure for a connection to a server
type server_conection struct {
	server_data    server
	created_at     time.Time
	conection_used int
}

func (s server) NewConectoin(password string) (*server_conection, error) {
	// Comparing the hash and the password
	var CheckHash string = s.hash
	Password_math, err := hash_password.CompareHash(password, CheckHash)

	// In case of error returning the error
	if err != nil {
		return nil, err
	}

	// In case passwords do not match returning an error
	if !Password_math {
		return nil, fmt.Errorf("401 Unauthorized: Passwords do not match")
	}

	// Returning a new server conection
	return &server_conection{server_data: s, created_at: time.Now(), conection_used: 0}, nil /* << It hasn't got and error so it's nill */
}
