//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fingerprint/configs"
	database "github.com/fingerprint/db"
	"github.com/fingerprint/dto"
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
	dto.NewValidator,
)

var HandlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewMinioHandler,
	handlers.NewOrganizationHandler,
	handlers.NewUserHandler,
	handlers.NewSiteHandler,
	handlers.NewCollectHandler,
	handlers.NewPointHandler,
)

var ServiceSet = wire.NewSet(
	services.NewAuthService,
	services.NewMinioService,
	services.NewCollectService,
)

var RepositorySet = wire.NewSet(
	repositories.NewMinioRepository,
	repositories.NewOrganizationRepository,
	repositories.NewUserRepository,
	repositories.NewSiteRepository,
	repositories.NewBuildingRepository,
	repositories.NewFloorRepository,
	repositories.NewPointRepository,
	repositories.NewCollectDeviceRepository,
	repositories.NewUploadRepository,
	repositories.NewFingerprintRepository,
)
