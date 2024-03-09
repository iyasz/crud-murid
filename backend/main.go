package main

import (
	"backend/controllers/classcontroller"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	models.ConnetDatabase()

	// Buat instansi Fiber
	app := fiber.New()

	app.Use(cors.New())

	// api class

	api := app.Group("/api")
	class := api.Group("/class")

	class.Get("/", classcontroller.Index)
	class.Post("/", classcontroller.Store)

	class.Get("/:id", classcontroller.Edit)
	class.Put("/:id", classcontroller.Update)
	class.Delete("/:id", classcontroller.Delete)

	// Atur rute dan handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Jalankan server dan dengarkan permintaan masuk di port 3000
	app.Listen(":8000")
}
