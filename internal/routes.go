package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/v1/shorten", GetShortenUrl)

}

func GetShortenUrl(c *fiber.Ctx) error {

	fmt.Println("wotking now ")

	return c.JSON("scueessfull")
}
