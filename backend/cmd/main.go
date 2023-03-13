package main

import (
	"backend/pkg/routes/example"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	example.Register(app)

	log.Fatal(app.Listen(":3000"))
}
