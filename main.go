package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/fingerprint/configs"
	_ "github.com/fingerprint/docs"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/fingerprint/routers"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func NewApp(
	middleware *middleware.AuthMiddleware,
	validator dto.Validator,
	authHandler handlers.AuthHandler,
	objectStorageHandler handlers.ObjectStorageHandler,
	objectStorageService services.ObjectStorageService,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
	siteHandler handlers.SiteHandler,
	localizeHandler handlers.LocalizeHandler,
	pointHandler handlers.PointHandler,
	buildingHandler handlers.BuildingHandler,
	floorHandler handlers.FloorHandler,
	trainingHandler handlers.MLHandler,
	dispatcherService services.DispatcherService,
) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.HandleError,
	})

	app.Use(
		logger.New(),
		middleware.ValidateJWT(),
		cors.New(cors.Config{AllowOrigins: "*"}),
	)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Set up routes
	routers.SetupRoutes(
		app.Group("/api"),
		validator,
		authHandler,
		objectStorageHandler,
		organizationHandler,
		userHandler,
		siteHandler,
		localizeHandler,
		pointHandler,
		trainingHandler,
		buildingHandler,
		floorHandler,
		middleware,
	)

	// Create default buckets
	objectStorageService.Initialize()

	// Start Queue
	dispatcherService.Start(context.Background())

	return app, nil
}

// @title Fingerprint API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
//
//go:generate swag init --parseDependency --parseInternal
func main() {

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
