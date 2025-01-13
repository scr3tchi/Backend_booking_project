package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scr3tchi/tiket-booking-project/config"
	"github.com/scr3tchi/tiket-booking-project/db"
	"github.com/scr3tchi/tiket-booking-project/handlers"
	"github.com/scr3tchi/tiket-booking-project/repositories"
)

func main() {

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigration)

	app := fiber.New(fiber.Config{
		AppName:      "tiket-booking",
		ServerHeader: "fiber",
	})

	//Repository
	eventRepository := repositories.NewEventRepository(db)

	//Router
	server := app.Group("/api")

	//Handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
