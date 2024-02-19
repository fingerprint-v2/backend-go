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
	configs.NewGRPCConnection,
	configs.NewGRPCClient,
	database.NewPostgresDatabase,
	middleware.NewAuthMiddleware,
	dto.NewValidator,
)

var HandlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewObjectStorageHandler,
	handlers.NewOrganizationHandler,
	handlers.NewUserHandler,
	handlers.NewSiteHandler,
	handlers.NewCollectHandler,
	handlers.NewPointHandler,
	handlers.NewMLHandler,
)

var ServiceSet = wire.NewSet(
	services.NewAuthService,
	services.NewObjectStorageService,
	services.NewCollectService,
	services.NewMLService,
	services.NewGRPCService,
)

var RepositorySet = wire.NewSet(
	repositories.NewOrganizationRepository,
	repositories.NewUserRepository,
	repositories.NewSiteRepository,
	repositories.NewBuildingRepository,
	repositories.NewFloorRepository,
	repositories.NewPointRepository,
	repositories.NewCollectDeviceRepository,
	repositories.NewUploadRepository,
	repositories.NewFingerprintRepository,
	repositories.NewWifiRepository,
)
