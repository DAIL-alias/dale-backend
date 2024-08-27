package repositories

import (
	"DALE/models"

	"gorm.io/gorm"
)

type AliasRepository struct {
	DB *gorm.DB
}

func NewAliasRepository(db *gorm.DB) *AliasRepository {
	return &AliasRepository{DB: db}
}

// Define Creating a new alias
func (r *AliasRepository) CreateAlias(alias *models.Alias) error {
	return r.DB.Create(alias).Error
}

// Define getting all aliases
func (r *AliasRepository) GetAliases() ([]models.Alias, error) {
	var aliases []models.Alias
	return aliases, r.DB.Find(&aliases).Error
}

// Define getting an alias by its id
func (r *AliasRepository) GetAliasByID(id int) (models.Alias, error) {
	var alias models.Alias
	return alias, r.DB.First(&alias, id).Error
}

func (r *AliasRepository) GetUsersAliases(userid int) ([]models.Alias, error) {
	var aliases []models.Alias
	err := r.DB.Debug().Where(`"user_id" = ? AND "is_deleted" = false`, userid).Find(&aliases).Error
	return aliases, err
}

func (r *AliasRepository) ToggleActiveStatus(id int) (models.Alias, error) {
	var alias models.Alias

	err := r.DB.First(&alias, id).Error
	if err != nil {
		return alias, err
	}

	alias.IsActive = !alias.IsActive

	err = r.DB.Save(&alias).Error
	return alias, err
}

// Delete alias
func (r *AliasRepository) DeleteAlias(id int) error {
	// Soft delete - set `IsDeleted` to true
	return r.DB.Model(&models.Alias{}).Where("id = ?", id).Update("is_deleted", true).Error
}