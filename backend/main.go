package main

import (
	"backend/controllers/classcontroller"
	"backend/controllers/muridcontroller"
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

	classApi := app.Group("/api")
	class := classApi.Group("/class")

	class.Get("/", classcontroller.Index)
	class.Post("/", classcontroller.Store)

	class.Get("/:id", classcontroller.Edit)
	class.Put("/:id", classcontroller.Update)
	class.Delete("/:id", classcontroller.Delete)

	// api murid

	muridApi := app.Group("/api")
	murid := muridApi.Group("/murid")

	murid.Get("/", muridcontroller.Index)
	murid.Post("/", muridcontroller.Store)
	murid.Get("/:id", muridcontroller.Edit)
	murid.Put("/:id", muridcontroller.Update)
	murid.Delete("/:id", muridcontroller.Delete)

	// Atur rute dan handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Jalankan server dan dengarkan permintaan masuk di port 3000
	app.Listen(":8000")
}
