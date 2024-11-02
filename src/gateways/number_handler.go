package gateways

import (
	"fiber-v3-template/src/domain/entities"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v3"
)

func (h *HTTPGateway) RandomNumber(c fiber.Ctx) error {
	data := h.NumberService.GenerateRandomNumber()
	return c.Status(200).JSON(entities.ResponseModel{
		Message: "Random number generated successfully",
		Data:    data,
	})
}

func (h *HTTPGateway) Calculate(c fiber.Ctx) error {
	var request entities.RequestModel
	payload := c.Body()
	if err := json.Unmarshal(payload, &request); err != nil {
		return c.Status(400).JSON(entities.ResponseModel{
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	result, err := h.NumberService.Calculate(request.A, request.B)
	if err != nil {
		return c.Status(400).JSON(entities.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(entities.ResponseModel{
		Message: "Calculated successfully",
		Data:    result,
	})

}
