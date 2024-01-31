package services

import (
	"github.com/fingerprint/repositories"
)

type OrganizationService interface {
	// SearchOrganization(ctx context.Context, org *validates.SearchOrganizationReq) ([]models.Organization, error)
}

type organizationServiceImpl struct {
	organizationRepo repositories.OrganizationRepository
}

func NewOrganizationService(organizationRepo repositories.OrganizationRepository) OrganizationService {
	return &organizationServiceImpl{
		organizationRepo: organizationRepo,
	}
}

// func (s *organizationServiceImpl) SearchOrganization(ctx context.Context, org *validates.SearchOrganizationReq) ([]models.Organization, error) {
// 	orgs, err := s.organizationRepo.SearchOrganization(ctx, org)
// 	if err != nil {
// 		return orgs, err
// 	}
// 	return orgs, nil
// }
