package server

import (
	"fmt"
	"time"
)

// Checks that the conection wasn't timed out or ha too many operations
func (s_c server_conection) CheckConection() (bool, error) {
	// Creating a date that if the max time of the conection
	nano_allowed := s_c.server_data.allowed_time * 1000000 * 1000
	max_time := s_c.created_at.Add(time.Duration(nano_allowed))

	// In case of the conection timeout returninig an error
	if max_time.Before(time.Now()) {
		return false, fmt.Errorf("408 Request Timeout: Server connection timed out")
	}

	// In case of too many operations with the conection  returninig an error
	if s_c.conection_used >= s_c.server_data.allowed_operations {
		return false, fmt.Errorf("429 Too Many Requests: Too many operations with server conection")
	}

	// Retuening true, that the conection is stil availible
	return true, nil
}

// Function to directly edit the server conection
func (s_c *server_conection) Touch() {
	// Increacing the conection_used
	s_c.conection_used++
}
