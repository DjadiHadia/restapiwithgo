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

func ListCars(c *fiber.Ctx) error {
	cars := []models.Car{}
	database.DB.Db.Find(&cars)

	return c.Status(200).JSON(cars)
}
