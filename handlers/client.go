package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
)

func AddClient(c *fiber.Ctx) error {
	client := new(models.Client)
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&client)

	return c.Status(200).JSON(client)
}

func ShowClientInfo(c *fiber.Ctx) error {
	client := models.Client{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&client)

	return c.Status(200).JSON(client)

}

func DeleteClient(c *fiber.Ctx) error {
	client := models.Client{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&client)
	if result.Error != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "client not found",
		})

	}

	return ListClients(c)
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}
	database.DB.Db.Find(&clients)

	return c.Status(200).JSON(clients)
}
