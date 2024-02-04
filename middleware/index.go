// middleware/withuser.go
package middleware

import (
	"context"
	"fiber-mongo-api/configs"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func WithUser(c *fiber.Ctx) error {
	user := c.Get("x-api-username")
	if user == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"error": "Missing x-api-username header (Middleware)"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := configs.GetCollection(configs.DB, "users")

	var existingUser bson.M
	err := userCollection.FindOne(ctx, bson.M{"userid": user}).Decode(&existingUser)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"error": "User not found in the database"})
	}

	return c.Next()
}
