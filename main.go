package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackcloudman/generic-go-api/routes"
)

func main() {
	app := fiber.New()
	// GET /john
	app.Get("/flights", routes.GetFlights)
	app.Post("/flights", routes.PostFlight)
	app.Get("/flights/:ID", routes.GetFlight)
	app.Delete("flights/:ID", routes.DeleteFlight)
	app.Put("flights/:ID", routes.PutFlight)
	log.Fatal(app.Listen(":10000"))
}
