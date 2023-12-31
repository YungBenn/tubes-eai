package routes

import (
	"Hotelin-BE/internal/controllers/admin"
	"Hotelin-BE/internal/controllers/hotel"
	"Hotelin-BE/internal/controllers/user"
	"Hotelin-BE/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// ==================== AUTH ========================

	api := app.Group("/api")

	register := api.Group("/register")
	register.Post("", user.Register)

	login := api.Group("/login")
	login.Post("", user.Login)

	// ==================== User ========================

	logoutUser := api.Group("/logoutUser").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	logoutUser.Post("", user.Logout)

	hotelAPI := api.Group("/hotels")
	hotelAPI.Get("", admin.ShowAllHotel)
	hotelAPI.Get("/details", hotel.GetHotelByID)

	userDetail := api.Group("/profile").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	userDetail.Get("", user.UserDetail)
	userDetail.Put("/change", user.ChangeProfile)

	balance := api.Group("/balance").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	balance.Post("/topup", user.TopUpBalance)
	balance.Get("", user.GetBalance)

	booking := api.Group("/booking").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	booking.Get("", user.ShowAllBooking)
	booking.Post("/create", user.CreateBooking)
	booking.Delete("/delete", user.DeleteBooking)

	room := api.Group("/rooms").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	room.Get("", hotel.GetRoomByHotelID)

	review := api.Group("/review").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	review.Post("/create", user.CreateReview)
	review.Get("", user.GetReview)

	// ================== End User =======================

	// =================== Admin =========================

	logoutAdmin := api.Group("/logoutAdmin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	logoutAdmin.Post("", user.Logout)

	adminAPI := api.Group("/admin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	adminAPI.Get("/hotel", admin.ShowAllHotel)

	// Consume Bundle
	adminAPI.Get("/bundle", admin.GetAllBundle)
	adminAPI.Get("/bundle/:id", admin.GetBundleByID)
	adminAPI.Put("/bundle/:id", admin.PutBookBundle)

	hotelAdminAPI := adminAPI.Group("/hotel")
	hotelAdminAPI.Post("/create", hotel.RegisterHotel)
	hotelAdminAPI.Get("/detail/:id", hotel.GetHotelByID)
	hotelAdminAPI.Put("/update/:id", hotel.UpdateHotel)
	hotelAdminAPI.Delete("/delete/:id", hotel.DeleteHotel)

	roomAPI := api.Group("/room").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	roomAPI.Post("/create", hotel.RegisterRoom)
	roomAPI.Put("/update", hotel.UpdateRoom)
	roomAPI.Delete("/delete", hotel.DeleteRoom)

}
