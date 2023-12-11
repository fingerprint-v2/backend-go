//go:build wireinject
// +build wireinject

package main

import (
	database "github.com/fingerprint/db"
	"github.com/fingerprint/handlers"
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
	database.NewPostgresDatabase,
)

var HandlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewOrganizationHandler,
	handlers.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	services.NewUserService,
	services.NewOrganizationService,
)

var RepositorySet = wire.NewSet(
	repositories.NewOrganizationRepository,
	repositories.NewUserRepository,
)
