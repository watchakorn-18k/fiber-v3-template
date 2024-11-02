package gateways

import (
	"fiber-v3-template/src/domain/entities"

	"github.com/gofiber/fiber/v3"
)

func (h *HTTPGateway) GetCarbon(c fiber.Ctx) error {
	url := c.Query("url")
	if url == "" {
		return c.Status(400).JSON(entities.ResponseMessage{
			Message: "Invalid request payload",
		})
	}
	carbon, err := h.CarbonProvider.GetCarbon(url)
	if err != nil {
		return c.Status(400).JSON(entities.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(entities.ResponseModel{
		Message: "Carbon calculated successfully",
		Data:    carbon,
	})
}
