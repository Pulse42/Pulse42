package main

import (
	"backend/pkg/routes/sample"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	sample.Register(app)

	log.Fatal(app.Listen(":3000"))
}
