package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddClient(c *fiber.Ctx) error {
	client := new(models.Client)
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	database.DB.Db.Create(&client)

	return c.Status(fiber.StatusCreated).JSON(client)
}

func ShowClientInfo(c *fiber.Ctx) error {
	client := models.Client{}
	id := c.Params("id")

	if err := database.DB.Db.Where("id = ?", id).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Client not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve client information",
		})
	}

	return c.Status(fiber.StatusOK).JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
	client := models.Client{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&client)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete client",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Client not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Client deleted successfully",
	})
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}
	if err := database.DB.Db.Find(&clients).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve clients",
		})
	}

	return c.Status(fiber.StatusOK).JSON(clients)
}

func UpdateClient(c *fiber.Ctx) error {
	client := new(models.Client)
	id := c.Params("id")

	// Parse the request body into the client object
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Check if the client with the given ID exists
	existingClient := models.Client{}
	if err := database.DB.Db.First(&existingClient, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Client not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve client information",
		})
	}

	// Update the existing client with the new values from the request body
	if client.Name != "" && client.Name != existingClient.Name {
		existingClient.Name = client.Name
	}
	if client.Email != "" && client.Email != existingClient.Email {
		existingClient.Email = client.Email
	}
	if client.Address != "" && client.Address != existingClient.Address {
		existingClient.Address = client.Address
	}
	if client.Phone != "" && client.Phone != existingClient.Phone {
		existingClient.Phone = client.Phone
	}
	if client.AgencyID != 0 && client.AgencyID != existingClient.AgencyID {
		existingClient.AgencyID = client.AgencyID
	}

	// Save the updated client to the database
	if err := database.DB.Db.Save(&existingClient).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update client",
		})
	}

	// Return the updated client
	return c.Status(fiber.StatusOK).JSON(existingClient)
}
