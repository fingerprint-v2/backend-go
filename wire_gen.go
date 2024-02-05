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
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userService)
	organizationRepository := repositories.NewOrganizationRepository(gormDB)
	authMiddleware := middleware.NewAuthMiddleware(authService, organizationRepository, userRepository)
	validator := dto.NewValidator()
	authHandler := handlers.NewAuthHandler(authService, userRepository)
	client := configs.NewMinioClient()
	minioRepository := repositories.NewMinioRepository(client)
	minioService := services.NewMinioService(minioRepository)
	minioHandler := handlers.NewMinioHandler(minioService)
	organizationService := services.NewOrganizationService(organizationRepository)
	organizationHandler := handlers.NewOrganizationHandler(organizationService, organizationRepository)
	userHandler := handlers.NewUserHandler(authService, userRepository, organizationRepository)
	siteRepository := repositories.NewSiteRepository(gormDB)
	siteHandler := handlers.NewSiteHandler(siteRepository)
	app, err := NewApp(authMiddleware, validator, authHandler, minioHandler, organizationHandler, userHandler, siteHandler)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
	}, nil
}

// wire.go:

var AppSet = wire.NewSet(
	NewApp, configs.NewMinioClient, db.NewPostgresDatabase, middleware.NewAuthMiddleware, dto.NewValidator,
)

var HandlerSet = wire.NewSet(handlers.NewAuthHandler, handlers.NewMinioHandler, handlers.NewOrganizationHandler, handlers.NewUserHandler, handlers.NewSiteHandler)

var ServiceSet = wire.NewSet(services.NewAuthService, services.NewMinioService, services.NewOrganizationService, services.NewUserService)

var RepositorySet = wire.NewSet(repositories.NewMinioRepository, repositories.NewOrganizationRepository, repositories.NewUserRepository, repositories.NewSiteRepository)
