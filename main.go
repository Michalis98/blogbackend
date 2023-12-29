package main

import (
	"log"
	"os"

	"github.com/Michalis98/blogbackend/database"
	"github.com/Michalis98/blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	//The package "github.com/gofiber/fiber/v2" is used for creating fast and flexible
	//HTTP servers in Go. Fiber is a web framework that is built on top of the fasthttp package,
	//which provides high-performance HTTP routing.
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
