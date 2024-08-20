package repositories

import (
	"gorm.io/gorm"
	"DALE/models"
)

// Struct defining methods for user DB operations
type UserRepository struct {
	DB *gorm.DB
}

// Creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create a new user from repository, optional error
func (r *UserRepository) CreateUser(user *models.User) (error) {
	return r.DB.Create(user).Error
}

// Get all users within the database
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

// Get user by ID
func (r *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return user, err
}