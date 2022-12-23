package main 

import "./database"
import "./routes"
import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/middleware/cors"

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

 routes.Setup(app)

	app.Listen(":8000")
}
