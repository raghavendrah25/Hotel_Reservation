package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrah25/Hotel_Reservation/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	listenAddress := flag.String("listen", ":8080", "The address to listen on for HTTP requests.")
	flag.Usage = func() {
		fmt.Println("Usage: server [options]")
		flag.PrintDefaults()
	}

	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddress)
}
