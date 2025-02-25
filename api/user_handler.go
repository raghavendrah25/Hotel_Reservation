package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrah25/Hotel_Reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Hello, User!",
	})
}
