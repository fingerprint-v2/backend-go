package db

import (
	"fmt"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrateModel(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Point{},
		&models.Floor{},
		&models.Building{},
		&models.Site{},
		&models.Organization{},
		&models.User{},
	); err != nil {
		return err
	}

	return nil
}

func NewPostgresDatabase() *gorm.DB {
	configs.InitialEnv(".env")
	configs := configs.GetPostgresConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", configs.Host, configs.User, configs.Password, configs.DBName, configs.Port, configs.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil
	}

	if err = migrateModel(db); err != nil {
		return nil
	}

	return db
}
