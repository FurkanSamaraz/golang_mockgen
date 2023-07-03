// controllers/api.go
package controllers

import (
	"strconv"

	"main/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) AddUser(ctx *fiber.Ctx) error {
	// Fiber API'si üzerinden gelen isteği işle
	// Gerekli parametreleri al
	name := ctx.FormValue("name")

	// UserService üzerinden AddUser işlemini çağır
	err := c.userService.AddUser(name)
	if err != nil {
		// Hata durumunda uygun yanıtı dön
		return err
	}

	// Başarılı yanıtı dön
	return ctx.JSON(fiber.Map{"message": "User added successfully"})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	// Fiber API'si üzerinden gelen isteği işle
	// Gerekli parametreleri al
	id := ctx.Params("id")
	ids, _ := strconv.Atoi(id)
	// ID'ye göre UserService üzerinden GetUserByID işlemini çağır
	user, err := c.userService.GetUserByID(ids)
	if err != nil {
		// Hata durumunda uygun yanıtı dön
		return err
	}

	// Başarılı yanıtı dön
	return ctx.JSON(fiber.Map{"user": user})
}
