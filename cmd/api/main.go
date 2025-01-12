package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scr3tchi/tiket-booking-project/handlers"
	"github.com/scr3tchi/tiket-booking-project/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "tiket-booking",
		ServerHeader: "fiber",
	})

	//Repository
	eventRepository := repositories.NewEventRepository(nil)

	//Router
	server := app.Group("/api")

	//Handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":9090")
}
