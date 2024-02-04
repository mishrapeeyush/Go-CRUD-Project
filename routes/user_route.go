package routes

import (
    "fiber-mongo-api/controllers"
    "github.com/gofiber/fiber/v2"
    "fiber-mongo-api/middleware"
)

func UserRoute(app *fiber.App) {
    router := app.Group("/api")

	router.Post("/user", controllers.CreateUser)
    router.Get("/user/:userId", middleware.WithUser, controllers.GetAUser)
    router.Put("/user/:userId", middleware.WithUser, controllers.EditAUser)
    router.Delete("/user/:userId", middleware.WithUser, controllers.DeleteAUser)
    router.Get("/users", middleware.WithUser, controllers.GetAllUsers)
}