// cmd/routes.go

package main

import (
	"github.com/DjadiHadia/restapiwithgo/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Apply authentication middleware to all routes

	// Apply authentication middleware to all routes except /login
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/login" || c.Path() == "/register" {
			return c.Next()
		}
		return handlers.AuthMiddleware(c)
	})
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)

	//-----------------agency routes----------------------
	app.Post("/agency", handlers.CreateAgency)

	app.Get("/agency/:id", handlers.ShowagencyInfo)
	app.Get("/agencies", handlers.ListAgencies)

	//-----------------car routes----------------------
	app.Get("/cars", handlers.ListCars)

	app.Post("/cars", handlers.AddCar)

	app.Get("/cars/:id", handlers.ShowCarInfo)

	app.Delete("/cars/:id", handlers.DeleteCar)

	app.Put("/cars/:id", handlers.UpdateCar)

	//-----------------client routes----------------------
	app.Get("/clients", handlers.ListClients)

	app.Post("/client", handlers.AddClient)

	app.Get("/client/:id", handlers.ShowClientInfo)

	app.Delete("/deleteclient/:id", handlers.DeleteClient)
	//-----------------reservation routes----------------------
	app.Get("/reservations", handlers.ListReservations)

	app.Post("/reservation", handlers.CreateReservation)

	app.Get("/reservation/:id", handlers.Showreservation)

	app.Delete("/deletereservation/:id", handlers.CancelReservation)

	//----------------user routes--------------------------
	app.Post("/register", handlers.RegisterUser)

	app.Post("/login", handlers.LoginUser)

}
