package repositories

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users *userRepository
}

func NewRepository(db *gorm.DB) Repo {
	return Repo{
		Users: newUserRepository(db),
	}

}
