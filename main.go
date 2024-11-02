package main

import (
	"fiber-v3-template/src/config"
	"fiber-v3-template/src/gateways"
	httpclient "fiber-v3-template/src/infra/http_client"
	"fiber-v3-template/src/middlewares"
	"fiber-v3-template/src/services"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	app := fiber.New(*config.NewConfig())
	middlewares.Logger(app)

	numberSV := services.NewNumberService()

	carbonProv := httpclient.NewCarbonProvider()

	gateways.NewHTTPGateway(app, numberSV, carbonProv)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(fmt.Sprintf(":%v", port))
}
