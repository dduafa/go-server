package repositories

import (
	"strings"

	"github.com/dduafa/go-server/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (rp *userRepository) Create(user *models.User) error {
	return rp.db.Create(user).Error
}

func (rp *userRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := rp.db.Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) FindByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) Update(user *models.User) error {
	return u.db.Save(user).Error
}

func (u *userRepository) Delete(id uuid.UUID) error {
	return u.db.Delete(&models.User{}, id).Error
}