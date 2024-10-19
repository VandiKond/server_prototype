package server

// You can add here some methods for the server (Please for safety use server conections func (s_c *server_conection))
// You can create operations with the database or with something like it

// A method to check all server parametrs
func (s_c *server_conection) ServerOk() (bool, error) {
	// Checking the conection
	ok, err := s_c.CheckConection()
	if err != nil {
		return false, err
	}

	// Increasing the server conection operations
	s_c.Touch()

	// Returning the statment of the server
	return ok, nil
}
