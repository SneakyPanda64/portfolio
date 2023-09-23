package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/api"
	"github.com/sneakypanda64/portfolio/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
	}
	app := fiber.New()

	db.Connect()

	api.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
