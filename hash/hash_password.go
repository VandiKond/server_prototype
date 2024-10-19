package hash_password

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

// A functionn that's getting the password and returning it's hash (sha-3-256)
func HashPassword(password string) (string, error) {
	// Creating a 256 byte hash as a slice of bytes
	hash := sha3.New256()
	_, err := hash.Write([]byte(password))

	// In case of error returning the error
	if err != nil {
		return "", fmt.Errorf("500 Internal Server Error: %e", err)
	}

	// Writing the slice in sha3
	sha3 := hash.Sum(nil)

	//Returninig the slice as a string
	return fmt.Sprintf("%x", sha3), nil /* << It hasn't got and error so it's nill */
}

// A function to compare hash of a password and the password
func CompareHash(password string, hash string) (bool, error) {
	// Creationg the hash of the password string
	hashedPassword, err := HashPassword(password)

	// In case of error returning the error
	if err != nil {
		return false, err
	}

	// Returning true if the hashs are the same. If not false
	return hashedPassword == hash, nil /* << It hasn't got and error so it's nill */
}
