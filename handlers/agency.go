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
	agency := models.Agency{}
	id := c.Params("Name")

	database.DB.Db.Where("Name = ?", id).First(&agency)

	return c.Status(200).JSON(agency)

}
