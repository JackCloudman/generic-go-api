package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackcloudman/generic-go-api/models"
)

type Response struct {
	Message string `json:"messsage"`
}

func GetFlights(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.Flights)
}

func GetFlight(c *fiber.Ctx) error {
	log.Printf("Flight: " + c.Params("ID"))
	return c.Status(fiber.StatusOK).JSON(models.Flights[0])
}
func PostFlight(c *fiber.Ctx) error {
	m := Response{
		Message: "Flight added (but not really)",
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

func DeleteFlight(c *fiber.Ctx) error {
	m := Response{
		Message: "Flight deleted (but not really)",
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

func PutFlight(c *fiber.Ctx) error {
	m := Response{
		Message: "Flight updated (but not really)",
	}
	return c.Status(fiber.StatusOK).JSON(m)
}
