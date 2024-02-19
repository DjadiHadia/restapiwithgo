// cmd/routes.go

package main

import (
	"github.com/DjadiHadia/restapiwithgo/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)
	//-----------------agency routes----------------------
	app.Post("/agency", handlers.CreateAgency)

	app.Get("/agency/:name", handlers.ShowagencyInfo)

	//-----------------car routes----------------------
	app.Get("/cars", handlers.ListCars)

	app.Post("/car", handlers.AddCar)

	app.Get("/car/:id", handlers.ShowCarInfo)

	app.Delete("/deletecar/:id", handlers.DeleteCar)

}
