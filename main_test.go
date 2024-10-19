package main

import (
	"go_server/server"
	"testing"
	"time"
)

var password string = "12345"
var max_operations int = 1
var serv, _ = server.NewServer(0, password, 1, max_operations)

func TestServer(t *testing.T) {

	t.Run("Error conection", func(t *testing.T) {
		_, err := serv.NewConectoin("SOME_OTHER_PASSWORD")
		if err == nil {
			t.Errorf("Error = %v; want %v", err, "401 Unauthorized: Passwords do not match")
		}
	})
	t.Run("Normal conection", func(t *testing.T) {
		_, err := serv.NewConectoin(password)
		if err != nil {
			t.Errorf("Error = %v; want %v", err, nil)
		}
	})
	t.Run("Too many operations", func(t *testing.T) {
		con, _ := serv.NewConectoin(password)
		ok, _ := con.ServerOk()
		if !ok {
			t.Errorf("Ok = %v; want %v", ok, true)
		}
		_, err := con.ServerOk()
		if err == nil {
			t.Errorf("Error = %v; want %v", err, "429 Too Many Requests: Too many operations with server conection")
		}
	})
	t.Run("Timeout", func(t *testing.T) {
		con, _ := serv.NewConectoin(password)
		time.Sleep(time.Second + time.Second)
		_, err := con.ServerOk()
		if err == nil {
			t.Errorf("Error = %v; want %v", err, "408 Request Timeout: Server connection timed out")
		}
	})

}
