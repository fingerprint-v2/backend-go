// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fingerprint/configs"
	"github.com/fingerprint/db"
	"github.com/fingerprint/handlers"
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
	authHandler := handlers.NewAuthHandler(userService)
	organizationRepository := repositories.NewOrganizationRepository(gormDB)
	organizationService := services.NewOrganizationService(organizationRepository)
	organizationHandler := handlers.NewOrganizationHandler(organizationService, organizationRepository)
	userHandler := handlers.NewUserHandler(userRepository)
	app, err := NewApp(authHandler, organizationHandler, userHandler)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
	}, nil
}

// wire.go:

var AppSet = wire.NewSet(
	NewApp, configs.NewMinioClient, db.NewPostgresDatabase,
)

var HandlerSet = wire.NewSet(handlers.NewAuthHandler, handlers.NewOrganizationHandler, handlers.NewUserHandler)

var ServiceSet = wire.NewSet(services.NewUserService, services.NewOrganizationService)

var RepositorySet = wire.NewSet(repositories.NewOrganizationRepository, repositories.NewUserRepository)
