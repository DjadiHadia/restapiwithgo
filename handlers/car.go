package handlers

import (
	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/gofiber/fiber/v2"
)

func AddCar(c *fiber.Ctx) error {
	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&car)

	return c.Status(200).JSON(car)
}

func ShowCarInfo(c *fiber.Ctx) error {
	car := models.Car{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&car)

	return c.Status(200).JSON(car)

}

func DeleteCar(c *fiber.Ctx) error {
	car := models.Car{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&car)
	if result.Error != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Car not found",
		})

	}

	return ListCars(c)
}

func UpdateCar(c *fiber.Ctx) error {
	car := new(models.Car)
	id := c.Params("id")

	// Parse the request body into the car object
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check if the car with the given ID exists
	existingCar := models.Car{}
	if err := database.DB.Db.First(&existingCar, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Car not found",
		})
	}

	// Update the existing car with the new values from the request body
	if car.Registration_number != "" {
		existingCar.Registration_number = car.Registration_number
	}
	if car.Brand != "" {
		existingCar.Brand = car.Brand
	}
	if car.Color != "" {
		existingCar.Color = car.Color
	}
	if car.Year != "" {
		existingCar.Year = car.Year
	}

	// If the agency_id is provided in the request body and not null, update it
	if car.AgencyID != 0 {
		existingCar.AgencyID = car.AgencyID
	}

	// Save the updated car to the database
	if err := database.DB.Db.Save(&existingCar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Return the updated car
	return c.Status(fiber.StatusOK).JSON(existingCar)
}

func ListCars(c *fiber.Ctx) error {
	cars := []models.Car{}
	database.DB.Db.Find(&cars)

	return c.Status(200).JSON(cars)
}
