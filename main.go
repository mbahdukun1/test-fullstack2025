package main

import (
	"go-fiber-login/config"
	"go-fiber-login/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	
	config.InitRedis()

	
	app.Post("/login", handlers.LoginHandler)

	app.Listen(":3030")
}
