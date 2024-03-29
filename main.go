package main

import (
    "fiber-mongo-api/configs"
    "fiber-mongo-api/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    //run database
    configs.ConnectDB()

    //connect to redis
    configs.InitRedis()

    //routes
    routes.UserRoute(app) 

    app.Listen(":1500")
}