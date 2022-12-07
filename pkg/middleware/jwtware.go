package middleware

import (
	"tobialbertino/portfolio-be/pkg/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var Cfg *config.Config

func init() {
	cfg, _ := config.LoadConfig()
	Cfg = cfg
}

func ProtectedJWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.GetKeyConfig("ACCESS_TOKEN_KEY")),
		SigningMethod: "HS256",
		AuthScheme:    "Bearer",
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
