// cmd/routes.go

package main

import (
	"github.com/DjadiHadia/restapiwithgo/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Apply authentication middleware to all routes
	app.Use(handlers.AuthMiddleware)
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)

	//-----------------agency routes----------------------
	app.Post("/agency", handlers.CreateAgency)

	app.Get("/agency/:name", handlers.ShowagencyInfo)
	app.Get("/agencies", handlers.ListAgencies)

	//-----------------car routes----------------------
	app.Get("/cars", handlers.ListCars)

	app.Post("/car", handlers.AddCar)

	app.Get("/car/:id", handlers.ShowCarInfo)

	app.Delete("/deletecar/:id", handlers.DeleteCar)

	app.Put("/cars/:id", handlers.UpdateCar)

	//-----------------client routes----------------------
	app.Get("/clients", handlers.ListClients)

	app.Post("/client", handlers.AddClient)

	app.Get("/client/:id", handlers.ShowClientInfo)

	app.Delete("/deleteclient/:id", handlers.DeleteClient)

	//----------------user routes--------------------------
	app.Post("/register", handlers.RegisterUser)

	app.Post("/login", handlers.LoginUser)

}
