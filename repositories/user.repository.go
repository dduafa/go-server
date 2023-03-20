package repositories

import (
	"github.com/dduafa/go-server/models"
	"strings"

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
	if err := rp.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (rp *userRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := rp.db.Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}
