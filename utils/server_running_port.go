package util

import "os"

type server struct{}

var ServerPort server

func (s *server) PortRunning(port string) string {
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = port
	}
	return serverPort
}
