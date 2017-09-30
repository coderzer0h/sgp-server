package main

import "github.com/coderzer0h/sgp-server/pkg/server"
import "github.com/Sirupsen/logrus"

func main() {
	server := server.NewServer()
	err := server.Start()
	if err != nil {
		logrus.Panicf("Server Crashed with error - %s", err)
	}
}
