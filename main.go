package main

import (
	"log"

	"main/controllers"
	repositorys "main/repository"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// UserRepository ve UserService mock nesnelerini oluştur
	userRepository := &repositorys.UserRepositoryImpl{}
	userService := services.NewUserService(userRepository)

	// UserController üzerinden Fiber API'larını işle
	userController := controllers.NewUserController(userService)

	app := fiber.New()

	app.Post("/users", userController.AddUser)
	app.Get("/users/:id", userController.GetUser)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
