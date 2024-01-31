package repositories

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/validates"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	repository[models.Organization, validates.SearchOrganizationReq]
	// SearchOrganization(context.Context, *validates.SearchOrganizationReq) ([]models.Organization, error)
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization, validates.SearchOrganizationReq]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Organization, validates.SearchOrganizationReq](db),
	}
}

// func (r organizationRepositoryImpl) SearchOrganization(ctx context.Context, org *validates.SearchOrganizationReq) ([]models.Organization, error) {
// 	orgs := []models.Organization{}
// 	// if err := r.db.Where("LOWER(name) LIKE LOWER(?)", org.Name+"%").Find(&orgs).Error; err != nil {
// 	// 	return nil, err
// 	// }

// 	// fmt.Println(org)

// 	orgJson, err := json.Marshal(org)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("------------------")
// 	fmt.Println(string(orgJson))
// 	fmt.Println("------------------")
// 	var orgMap map[string]interface{}

// 	json.Unmarshal(orgJson, &orgMap)

// 	for field, val := range orgMap {
// 		fmt.Println("KV Pair: ", field, val)
// 	}

// 	// fmt.Println(string(orgJson))
// 	if err := r.db.Find(&orgs, orgMap).Error; err != nil {
// 		return nil, err
// 	}

// 	return orgs, nil
// }
