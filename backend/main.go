package main

import (
	"log"

	"github.com/JerryJeager/JeagerEats/cmd"
	"github.com/JerryJeager/JeagerEats/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	config.ConnectToRedis()
}

func main() {
	log.Println("Starting JeagerEats Server")
	cmd.ExecuteApiRoutes()
}
