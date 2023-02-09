package sample

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	sample := app.Group("/sample")
	sample.Get("/hello", helloHandler)
	sample.Get("/world", worldHandler)
}
