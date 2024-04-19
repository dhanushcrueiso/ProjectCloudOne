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

var URLMap = make(map[string]string)
var basepath = "http://shortn.ly/"

func GetShortenUrl(c *fiber.Ctx) error {

	req := RequestData{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}
	fmt.Println("check 1")
	originalURL := req.URL
	if _, exists := URLMap[originalURL]; !exists {
		// Generate shortened URL
		fmt.Println("check 2")
		shortenedURL := generateShortURL()
		URLMap[originalURL] = basepath + shortenedURL
	}
	fmt.Println("check 3")
	res := RequestData{
		URL: URLMap[originalURL],
	}
	return c.JSON(res)
}

func Redirect(c *fiber.Ctx) error {

	req := RequestData{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}
	fmt.Println("check 1")

	res := getKeyByValue(URLMap, req.URL)

	return c.JSON(res)

}
