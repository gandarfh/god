package main

import (
	"log"
	"net/http"

	"github.com/gandarfh/god"
	"github.com/gofiber/fiber/v2"
)

var bodySchema = god.Object(god.Map{
	"name":     god.String(god.Required()),
	"lastName": god.String(),
	"email":    god.String(god.Required(), god.Email()),
})

func HelloMapController(c *fiber.Ctx) error {
	body := map[string]interface{}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]error{
			"error": err,
		})
	}

	if err := god.Validate(body, bodySchema); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(god.ErrorsToMap(err))
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"message": "sucess",
	})
}

type Payload struct {
	Email    string `query:"email"`
	Name     string `json:"name"`
	LastName string `god:"lastName"`
}

func HelloStructController(c *fiber.Ctx) error {
	body := Payload{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err := god.Validate(body, bodySchema); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(god.ErrorsToMap(err))
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"message": "sucess",
	})
}

func main() {
	app := fiber.New()
	app.Post("/hello/map", HelloMapController)
	app.Post("/hello/struct", HelloStructController)

	log.Fatal(app.Listen(":3000"))
}
