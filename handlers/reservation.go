package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
)

func CreateReservation(c *fiber.Ctx) error {
	reservation := new(models.Reservation)
	if err := c.BodyParser(reservation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&reservation)

	return c.Status(200).JSON(reservation)
}
func ListReservations(c *fiber.Ctx) error {
	reservations := []models.Reservation{}
	database.DB.Db.Find(&reservations)

	return c.Status(200).JSON(reservations)
}
func Showreservation(c *fiber.Ctx) error {
	reservation := models.Reservation{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&reservation)

	return c.Status(200).JSON(reservation)

}
func CancelReservation(c *fiber.Ctx) error {
	reservation := models.Reservation{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&reservation)
	if result.Error != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "reservation not found",
		})

	}

	return ListReservations(c)
}
