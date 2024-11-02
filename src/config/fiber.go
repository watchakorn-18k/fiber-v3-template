package config

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

func NewConfig() *fiber.Config {
	return &fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ServerHeader: "fiber-v3-template",
		AppName:      "fiber-v3-template",
	}
}
