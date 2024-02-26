package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReservation(c *fiber.Ctx) error {
	reservation := new(models.Reservation)
	if err := c.BodyParser(reservation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to parse request body",
		})
	}

	database.DB.Db.Create(&reservation)

	return c.Status(fiber.StatusCreated).JSON(reservation)
}

func ListReservations(c *fiber.Ctx) error {
	reservations := []models.Reservation{}
	database.DB.Db.Find(&reservations)

	return c.Status(fiber.StatusOK).JSON(reservations)
}

func Showreservation(c *fiber.Ctx) error {
	reservation := models.Reservation{}
	id := c.Params("id")

	if err := database.DB.Db.Where("id = ?", id).First(&reservation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Reservation not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve reservation information",
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservation)
}

func CancelReservation(c *fiber.Ctx) error {
	id := c.Params("id")

	reservation := models.Reservation{}
	result := database.DB.Db.Where("id = ?", id).Delete(&reservation)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete reservation",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Reservation not found",
		})
	}

	return ListReservations(c)
}
