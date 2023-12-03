package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/routers"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewApp(organizationHandler handlers.OrganizationHandler, userHandler handlers.UserHandler) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.HanndleError,
	})

	app.Use(
		logger.New(),
		cors.New(cors.Config{AllowOrigins: "*"}),
	)

	// Set up routes
	routers.SetupRoutes(
		app.Group("/api"),
		organizationHandler,
		userHandler,
	)

	return app, nil
}

//go:generate swag init
func main() {

	configs.InitialEnv(".env")

	app, cleanup, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)
	go func() {
		<-exitChan
		cleanup()
		app.Shutdown()
		os.Exit(0)
	}()

	if err := app.Listen(fmt.Sprintf(":%s", configs.GetPort())); err != nil {
		log.Fatal(err)
	}
}
