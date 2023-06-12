package hotel

import (
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterHotel(c *fiber.Ctx) error {
	req := models.HotelRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	hotel := models.Hotel{
		Name:       req.Name,
		Email:			req.Email,
		Phone:			req.Phone,
		Address:		req.Address,
		Rating:			req.Rating,
	}

	err := database.DB.Create(&hotel).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Hotel Added",
	})
}


func GetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var hotel =  models.Hotel{ID: id}

	// err := database.DB.Where("id = ?", id).First(&hotel).Error
	err := database.DB.Model([]models.Hotel{}).First(&hotel).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hotelResponse := models.HotelResponse{
		ID:      hotel.ID,
		Name:    hotel.Name,
		Email:   hotel.Email,
		Phone:   hotel.Phone,
		Address: hotel.Address,
		Rating:	 hotel.Rating,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hotel found",
		"data":    hotelResponse,
	})

}

func UpdateHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	// var hotel models.Hotel

	req := models.HotelRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	updateHotel := models.Hotel{
		ID: string(id),
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Address: req.Address,
		Rating: req.Rating,
	}

	err := database.DB.Save(&updateHotel).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hotel updated",
	})
}

func DeleteHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	// var hotel models.Hotel
	hotel := models.Hotel{
		ID: string(id),
	}

	err := database.DB.Delete(&hotel).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hotel deleted",
	})
}