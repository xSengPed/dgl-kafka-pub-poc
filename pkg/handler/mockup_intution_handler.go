package handler

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

type MockupIntutionHandler struct {
}

func NewMockupIntutionHandler() *MockupIntutionHandler {
	return &MockupIntutionHandler{}
}

func (h *MockupIntutionHandler) PostToInstution(ctx *fiber.Ctx) error {

	data, err := os.ReadFile("mockup/resp_intuition.json")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	var text interface{}
	err = json.Unmarshal(data, &text)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(text)
}
