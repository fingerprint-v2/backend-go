package main

import (
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fingerprint/db"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"gorm.io/gorm"
)

type Seeder interface {
	seed()
	seedOrganization() []models.Organization
	seedUser([]models.Organization) []models.User
}

type seederImpl struct {
	faker       *gofakeit.Faker
	db          *gorm.DB
	authService services.AuthService
}

func NewSeeder(faker *gofakeit.Faker, db *gorm.DB, authService services.AuthService) Seeder {
	return &seederImpl{
		faker:       faker,
		db:          db,
		authService: authService,
	}
}

func (s *seederImpl) seed() {
	orgs := s.seedOrganization()
	s.seedUser(orgs)
}

func (s *seederImpl) seedOrganization() []models.Organization {
	var orgs []models.Organization
	for i := 0; i < 1; i++ {
		org := models.Organization{
			Name: s.faker.Company(),
		}
		orgs = append(orgs, org)
	}

	if err := s.db.Create(&orgs).Error; err != nil {
		return nil
	}

	return orgs
}

func (s *seederImpl) seedUser(orgs []models.Organization) []models.User {
	var users []models.User
	var wg sync.WaitGroup
	hashedUsers := make(chan models.User, 100)

	for i := 0; i < 2; i++ {
		user := models.User{
			Username:       s.faker.Username(),
			Password:       "1234",
			Role:           "USER",
			OrganizationID: orgs[0].ID.String(),
		}

		if i == 1 {
			user.Role = "SUPER_ADMIN"
		}

		wg.Add(1)
		go func(u models.User) {
			defer wg.Done()
			s.authService.HashPassword(&u)
			hashedUsers <- u
		}(user)
	}

	go func() {
		wg.Wait()
		close(hashedUsers)
	}()

	for hashedUser := range hashedUsers {
		users = append(users, hashedUser)
	}

	if err := s.db.Create(&users).Error; err != nil {
		return nil
	}

	return users
}

func main() {
	faker := gofakeit.New(0)
	db := db.NewPostgresDatabase()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userService)
	NewSeeder(faker, db, authService).seed()
}
