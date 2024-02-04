package main

import (
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fingerprint/constants"
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
	s.resetDB()
	orgs := s.seedOrganization()
	s.seedUser(orgs)
	s.seedSite(orgs)
}

func (s *seederImpl) seedOrganization() []models.Organization {
	var orgs []models.Organization
	for i := 0; i < 3; i++ {
		org := models.Organization{
			Name: s.faker.Company(),
		}
		if i == 0 {
			org.Name = "SuperAdminOrg"
		}
		if i == 1 {
			org.Name = "org1"
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

	for i := 0; i < 5; i++ {
		user := models.User{
			Username:       s.faker.Username(),
			Password:       "1234",
			Role:           constants.USER.String(),
			OrganizationID: orgs[1].ID.String(),
		}

		if i == 1 {
			user.Role = constants.SUPERADMIN.String()
			user.Username = "superadmin1"
			user.OrganizationID = orgs[0].ID.String()
		}

		if i == 2 {
			user.Role = constants.ADMIN.String()
			user.Username = "admin1"
		}

		if i == 3 {
			user.Role = constants.USER.String()
			user.Username = "user1"
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

func (s *seederImpl) seedSite(orgs []models.Organization) []models.Site {

	var sites []models.Site
	for _, org := range orgs[1:] {
		for siteIdx := 0; siteIdx < 1; siteIdx++ {
			site := models.Site{
				Name:           s.faker.Name(),
				OrganizationID: org.ID.String(),
			}
			sites = append(sites, site)
		}
	}

	sites[0].Name = "site1"

	if err := s.db.Create(&sites).Error; err != nil {
		return nil
	}

	return sites
}

func (s *seederImpl) resetDB() {
	s.db.Exec("TRUNCATE TABLE users CASCADE")
	s.db.Exec("TRUNCATE TABLE sites CASCADE")
	s.db.Exec("TRUNCATE TABLE organizations CASCADE")
}

func main() {
	faker := gofakeit.New(0)
	db := db.NewPostgresDatabase()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userService)
	NewSeeder(faker, db, authService).seed()
}
