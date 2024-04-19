package main

import (
	"infracloud/config"
	"infracloud/internal"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	env := "dev"

	var file *os.File
	var err error

	file, err = os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}

	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)

	app := fiber.New()

	internal.SetupRoutes(app)

	// r := internal.GetRouter()
	// constants.Logger.Info("Listening to Port: " + cnf.Port)
	// r.Run(":" + cnf.Port)
	app.Listen(cnf.Port)
}
