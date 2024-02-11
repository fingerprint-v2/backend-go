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
	seedSite([]models.Organization) []models.Site
	seedBuilding([]models.Site) []models.Building
	seedFloor([]models.Building) []models.Floor
	seedPoint([]models.Floor) []models.Point
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
	sites := s.seedSite(orgs)
	buildings := s.seedBuilding(sites)
	floors := s.seedFloor(buildings)
	s.seedPoint(floors)
}

func (s *seederImpl) seedOrganization() []models.Organization {
	var orgs []models.Organization
	for i := 0; i < 10; i++ {
		org := models.Organization{
			Name: s.faker.Company(),
		}
		orgs = append(orgs, org)
	}

	orgs[0].Name = "SuperAdminOrg"
	orgs[0].IsSystem = true
	orgs[1].Name = "org1"

	if err := s.db.Create(&orgs).Error; err != nil {
		return nil
	}

	return orgs
}

func (s *seederImpl) seedUser(orgs []models.Organization) []models.User {
	var users []models.User
	var wg sync.WaitGroup
	hashedUsers := make(chan models.User, 100)

	// Superadmin1
	user := models.User{
		Username:       "superadmin1",
		Password:       "1234",
		Role:           constants.SUPERADMIN.String(),
		OrganizationID: orgs[0].ID.String(),
	}
	users = append(users, user)

	// Admin1
	user = models.User{
		Username:       "admin1",
		Password:       "1234",
		Role:           constants.ADMIN.String(),
		OrganizationID: orgs[1].ID.String(),
	}
	users = append(users, user)

	// User1
	user = models.User{
		Username:       "user1",
		Password:       "1234",
		Role:           constants.USER.String(),
		OrganizationID: orgs[1].ID.String(),
	}
	users = append(users, user)

	// Rest
	for _, org := range orgs[2:] {
		for i := 0; i < 2; i++ {
			user := models.User{
				Username:       s.faker.Username(),
				Password:       "1234",
				Role:           constants.USER.String(),
				OrganizationID: org.ID.String(),
			}
			users = append(users, user)
		}
	}

	for _, user := range users {
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

	// Resetting
	users = []models.User{}
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
		for siteIdx := 0; siteIdx < 2; siteIdx++ {
			site := models.Site{
				Name:           "site_" + s.faker.Word(),
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

func (s *seederImpl) seedBuilding(sites []models.Site) []models.Building {

	var buildings []models.Building
	for _, site := range sites {
		for buildingIdx := 0; buildingIdx < 2; buildingIdx++ {
			building := models.Building{
				Name:           "building_" + s.faker.Word(),
				SiteID:         site.ID.String(),
				OrganizationID: site.OrganizationID,
			}
			buildings = append(buildings, building)
		}
	}

	buildings[0].Name = "building1"

	if err := s.db.Create(&buildings).Error; err != nil {
		return nil
	}

	return buildings
}

func (s *seederImpl) seedFloor(buildings []models.Building) []models.Floor {

	var floors []models.Floor
	for _, building := range buildings {
		for floorIdx := 0; floorIdx < 2; floorIdx++ {
			floor := models.Floor{
				Name:           "floor_" + s.faker.Word(),
				BuildingID:     building.ID.String(),
				SiteID:         building.SiteID,
				OrganizationID: building.OrganizationID,
			}
			floors = append(floors, floor)
		}
	}

	floors[0].Name = "floor1"

	if err := s.db.Create(&floors).Error; err != nil {
		return nil
	}

	return floors
}

func (s *seederImpl) seedPoint(floors []models.Floor) []models.Point {

	var points []models.Point
	for _, floor := range floors {
		for pointIdx := 0; pointIdx < 2; pointIdx++ {
			point := models.Point{
				Name:           "point_" + s.faker.Word(),
				FloorID:        floor.ID.String(),
				BuildingID:     floor.BuildingID,
				SiteID:         floor.SiteID,
				OrganizationID: floor.OrganizationID,
			}
			points = append(points, point)
		}
	}

	points[0].Name = "point1"

	if err := s.db.Create(&points).Error; err != nil {
		return nil
	}

	return points
}

func (s *seederImpl) resetDB() {
	s.db.Exec("TRUNCATE TABLE points CASCADE")
	s.db.Exec("TRUNCATE TABLE floors CASCADE")
	s.db.Exec("TRUNCATE TABLE buildings CASCADE")
	s.db.Exec("TRUNCATE TABLE sites CASCADE")
	s.db.Exec("TRUNCATE TABLE users CASCADE")
	s.db.Exec("TRUNCATE TABLE organizations CASCADE")
}

func main() {
	faker := gofakeit.New(0)
	db := db.NewPostgresDatabase()
	//
	userRepo := repositories.NewUserRepository(db)
	siteRepo := repositories.NewSiteRepository(db)
	buildingRepo := repositories.NewBuildingRepository(db)
	floorRepo := repositories.NewFloorRepository(db)
	pointRepo := repositories.NewPointRepository(db)
	//
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userService, userRepo, siteRepo, buildingRepo, floorRepo, pointRepo)
	NewSeeder(faker, db, authService).seed()
}
