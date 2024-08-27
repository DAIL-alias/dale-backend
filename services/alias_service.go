package services

import (
	"DALE/models"
	"DALE/repositories"
)

type AliasService struct {
	AliasRepository *repositories.AliasRepository
}

func NewAliasService(aliasRepository *repositories.AliasRepository) *AliasService {
	return &AliasService{AliasRepository: aliasRepository}
}

func (s *AliasService) CreateAlias(alias *models.Alias) error {
	return s.AliasRepository.CreateAlias(alias)
}

func (s *AliasService) GetAliases() ([]models.Alias, error) {
	return s.AliasRepository.GetAliases()
}

func (s *AliasService) GetAliasByID(id int) (models.Alias, error) {
	return s.AliasRepository.GetAliasByID(id)
}

func (s *AliasService) GetUsersAliases(userID int) ([]models.Alias, error) {
	return s.AliasRepository.GetUsersAliases(userID)
}

func (s *AliasService) ToggleActiveStatus(id int) (models.Alias, error) {
	return s.AliasRepository.ToggleActiveStatus(id)
}

func (s *AliasService) DeleteAlias(id int) error {
	return s.AliasRepository.DeleteAlias(id)
}