package main

import "github.com/GoRedTeam/Realm/internal/pkg/server"

func main() {

	// Set up server
	server.Listen(server.Init())

}
