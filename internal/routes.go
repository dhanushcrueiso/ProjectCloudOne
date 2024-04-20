package internal

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/v1/shorten", GetShortenUrl)
	app.Post("/v1/getOriginal", Redirect)

}

type RequestData struct {
	URL string `json:"url"`
}

func GetShortenUrl(c *fiber.Ctx) error {

	req := RequestData{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	res := ShorternUrl(req)

	return c.JSON(res)
}

func Redirect(c *fiber.Ctx) error {

	req := RequestData{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}
	fmt.Println("check 1")

	res := getKeyByValue(req)
	resjson := RequestData{
		URL: res,
	}
	return c.JSON(resjson)

}
