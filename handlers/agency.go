package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAgency(c *fiber.Ctx) error {
	agency := new(models.Agency)
	if err := c.BodyParser(agency); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&agency)

	return c.Status(200).JSON(agency)
}

func ShowagencyInfo(c *fiber.Ctx) error {
	// Retrieve agency ID from request parameters
	id := c.Params("id")

	// Fetch agency from the database including its associated cars
	var agency models.Agency
	if err := database.DB.Db.Preload("Cars").Where("id = ?", id).First(&agency).Error; err != nil {
		// Handle error if agency not found or any other database error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Agency not found",
		})
	}

	// Return agency data along with associated cars
	return c.Status(fiber.StatusOK).JSON(agency)
}

func ListAgencies(c *fiber.Ctx) error {
	agencies := []models.Agency{}

	database.DB.Db.Find(&agencies)

	return c.Status(200).JSON(agencies)

}
