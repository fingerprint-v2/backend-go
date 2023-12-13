//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fingerprint/configs"
	database "github.com/fingerprint/db"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeApp() (*fiber.App, func(), error) {
	wire.Build(AppSet, HandlerSet, ServiceSet, RepositorySet)

	return &fiber.App{}, func() {}, nil
}

var AppSet = wire.NewSet(
	NewApp,
	configs.NewMinioClient,
	database.NewPostgresDatabase,
	middleware.NewAuthMiddleware,
)

var HandlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewOrganizationHandler,
	handlers.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	services.NewAuthService,
	services.NewUserService,
	services.NewOrganizationService,
)

var RepositorySet = wire.NewSet(
	repositories.NewOrganizationRepository,
	repositories.NewUserRepository,
)
