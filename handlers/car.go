package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddCar(c *fiber.Ctx) error {
	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	if err := database.DB.Db.Create(car).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to add car",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(car)
}

func ShowCarInfo(c *fiber.Ctx) error {
	car := models.Car{}
	id := c.Params("id")

	if err := database.DB.Db.Where("id = ?", id).First(&car).Error; err != nil {
		if database.DB.Db.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Car not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve car information",
		})
	}

	return c.Status(fiber.StatusOK).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Db.Delete(&models.Car{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete car",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Car not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Car deleted successfully",
	})
}

func UpdateCar(c *fiber.Ctx) error {
	car := new(models.Car)
	id := c.Params("id")

	// Parse the request body into the car object
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Check if the car with the given ID exists
	existingCar := models.Car{}
	if err := database.DB.Db.First(&existingCar, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Car not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve car information",
		})
	}

	// Update the existing car with the new values from the request body
	if car.Registration_number != "" && car.Registration_number != existingCar.Registration_number {
		existingCar.Registration_number = car.Registration_number
	}
	if car.Brand != "" && car.Brand != existingCar.Brand {
		existingCar.Brand = car.Brand
	}
	if car.Color != "" && car.Color != existingCar.Color {
		existingCar.Color = car.Color
	}
	if car.Year != "" && car.Year != existingCar.Year {
		existingCar.Year = car.Year
	}

	// If the agency_id is provided in the request body and not null, update it
	if car.AgencyID != 0 && car.AgencyID != existingCar.AgencyID {
		existingCar.AgencyID = car.AgencyID
	}

	// Save the updated car to the database
	if err := database.DB.Db.Save(&existingCar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update car",
		})
	}

	// Return the updated car
	return c.Status(fiber.StatusOK).JSON(existingCar)
}

func ListCars(c *fiber.Ctx) error {
	cars := []models.Car{}

	if err := database.DB.Db.Find(&cars).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to retrieve cars",
		})
	}

	return c.Status(fiber.StatusOK).JSON(cars)
}
