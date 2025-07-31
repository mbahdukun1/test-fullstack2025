package handlers

import (
	"encoding/json"
	"go-fiber-login/config"
	"go-fiber-login/models"
	"go-fiber-login/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	key := "login_" + input.Username
	val, err := config.Redis.Get(config.Ctx, key).Result()
	if err == redis.Nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Redis error"})
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Parse error"})
	}

	if user.Password != utils.Sha1Hash(input.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Wrong password"})
	}

	return c.JSON(fiber.Map{
		"message":  "Login success",
		"realname": user.RealName,
		"email":    user.Email,
	})
}
