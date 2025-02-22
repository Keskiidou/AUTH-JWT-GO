package main

import (
	"GO_Auth/initializers"
	"GO_Auth/server"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	server.StartServer()
}
