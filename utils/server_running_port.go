package app

import "os"

type server struct{}

var Server server

func (s *server) ServerRunningPort(port string) string {
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = port
	}
	return serverPort
}
