package main

import (
	"log"
	"regex-engine/pkg/regex"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	log.Println("Server Started")
	app := fiber.New()

	// Enable CORS with default settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Post("/api/match", func(c *fiber.Ctx) error {
		var request struct {
			Pattern string `json:"pattern"`
			Text    string `json:"text"`
		}

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		log.Printf("Patter: %s, Text: %s\n", request.Pattern, request.Text)

		matched, matches := regex.Match(request.Pattern, request.Text)
		log.Printf("Match?: %t, %v\n", matched, matches)

		return c.JSON(fiber.Map{
			"matched": matched,
			"matches": matches,
		})
	})

	//start the server
	log.Fatal(app.Listen(":3100"))
}
