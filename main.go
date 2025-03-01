package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrah25/Hotel_Reservation/api"
	"github.com/raghavendrah25/Hotel_Reservation/types"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.Usage = func() {
		fmt.Println("Usage: server [options]")
		flag.PrintDefaults()
	}

	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(":" + port)
}
