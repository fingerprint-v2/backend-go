// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fingerprint/configs"
	"github.com/fingerprint/db"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/middlewares"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

import (
	_ "github.com/fingerprint/docs"
)

// Injectors from wire.go:

func InitializeApp() (*fiber.App, func(), error) {
	gormDB := db.NewPostgresDatabase()
	userRepository := repositories.NewUserRepository(gormDB)
	siteRepository := repositories.NewSiteRepository(gormDB)
	buildingRepository := repositories.NewBuildingRepository(gormDB)
	floorRepository := repositories.NewFloorRepository(gormDB)
	pointRepository := repositories.NewPointRepository(gormDB)
	authService := services.NewAuthService(userRepository, siteRepository, buildingRepository, floorRepository, pointRepository)
	organizationRepository := repositories.NewOrganizationRepository(gormDB)
	authMiddleware := middleware.NewAuthMiddleware(authService, organizationRepository, userRepository)
	validator := dto.NewValidator()
	authHandler := handlers.NewAuthHandler(authService, userRepository)
	client := configs.NewMinioClient()
	objectStorageService := services.NewObjectStorageService(client)
	objectStorageHandler := handlers.NewObjectStorageHandler(objectStorageService)
	organizationHandler := handlers.NewOrganizationHandler(organizationRepository)
	userHandler := handlers.NewUserHandler(authService, userRepository, organizationRepository)
	siteHandler := handlers.NewSiteHandler(siteRepository)
	collectDeviceRepository := repositories.NewCollectDeviceRepository(gormDB)
	uploadRepository := repositories.NewUploadRepository(gormDB)
	fingerprintRepository := repositories.NewFingerprintRepository(gormDB)
	wifiRepository := repositories.NewWifiRepository(gormDB)
	collectService := services.NewCollectService(collectDeviceRepository, uploadRepository, fingerprintRepository, wifiRepository, pointRepository, siteRepository)
	collectHandler := handlers.NewCollectHandler(collectService)
	pointHandler := handlers.NewPointHandler(pointRepository)
	clientConn, cleanup, err := configs.NewGRPCConnection()
	if err != nil {
		return nil, nil, err
	}
	fingperintClient := configs.NewGRPCClient(clientConn)
	grpcService := services.NewGRPCService(fingperintClient)
	mlService := services.NewMLService(objectStorageService, grpcService)
	mlHandler := handlers.NewMLHandler(mlService)
	app, err := NewApp(authMiddleware, validator, authHandler, objectStorageHandler, objectStorageService, organizationHandler, userHandler, siteHandler, collectHandler, pointHandler, mlHandler)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

var AppSet = wire.NewSet(
	NewApp, configs.NewMinioClient, configs.NewGRPCConnection, configs.NewGRPCClient, db.NewPostgresDatabase, middleware.NewAuthMiddleware, dto.NewValidator,
)

var HandlerSet = wire.NewSet(handlers.NewAuthHandler, handlers.NewObjectStorageHandler, handlers.NewOrganizationHandler, handlers.NewUserHandler, handlers.NewSiteHandler, handlers.NewCollectHandler, handlers.NewPointHandler, handlers.NewMLHandler)

var ServiceSet = wire.NewSet(services.NewAuthService, services.NewObjectStorageService, services.NewCollectService, services.NewMLService, services.NewGRPCService)

var RepositorySet = wire.NewSet(repositories.NewOrganizationRepository, repositories.NewUserRepository, repositories.NewSiteRepository, repositories.NewBuildingRepository, repositories.NewFloorRepository, repositories.NewPointRepository, repositories.NewCollectDeviceRepository, repositories.NewUploadRepository, repositories.NewFingerprintRepository, repositories.NewWifiRepository)
