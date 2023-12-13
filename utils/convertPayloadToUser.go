package utils

import (
	"errors"

	"github.com/fingerprint/models"
	"github.com/google/uuid"
)

func ConvertPayloadToUser(payload map[string]interface{}) (*models.User, error) {
	idString, idExists := payload["id"].(string)
	if !idExists {
		return nil, errors.New("user ID is missing or not a string")
	}
	userId, err := uuid.Parse(idString)
	if err != nil {
		return nil, errors.New("failed to parse user ID: " + idString)
	}

	user := &models.User{
		ID:             userId,
		Username:       payload["username"].(string),
		Role:           payload["role"].(string),
		OrganizationID: payload["organization_id"].(string),
	}

	return user, nil

}
