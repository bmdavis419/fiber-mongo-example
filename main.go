package main

import (
	"os"

	"github.com/bmdavis419/fiber-mongo-example/common"
	"github.com/bmdavis419/fiber-mongo-example/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	// init env
	err := common.LoadEnv()
	if err != nil {
		return err
	}

	// init db
	err = common.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	defer common.CloseDB()

	// create app
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// add routes
	router.AddBookGroup(app)

	// start server
	var port string
	if port = os.Getenv("PORT"); port != "" {
		port = "8080"
	}
	app.Listen(":" + port)

	return nil
}
