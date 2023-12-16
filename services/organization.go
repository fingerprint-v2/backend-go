package services

import (
	"context"

	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
)

type OrganizationService interface {
	SearchOrganization(ctx context.Context, org *models.Organization) ([]models.Organization, error)
}

type organizationServiceImpl struct {
	organizationRepo repositories.OrganizationRepository
}

func NewOrganizationService(organizationRepo repositories.OrganizationRepository) OrganizationService {
	return &organizationServiceImpl{
		organizationRepo: organizationRepo,
	}
}

func (s *organizationServiceImpl) SearchOrganization(ctx context.Context, org *models.Organization) ([]models.Organization, error) {
	orgs, err := s.organizationRepo.SearchOrganization(ctx, org)
	if err != nil {
		return orgs, err
	}
	return orgs, nil
}
