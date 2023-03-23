package services

import (
	"github.com/dduafa/go-server/core"
	"github.com/dduafa/go-server/models"
	"github.com/dduafa/go-server/repositories"
	"github.com/google/uuid"
)

type userService struct {
	repository repositories.Repo
	config     *core.Config
}

func newUserService(r repositories.Repo, c *core.Config) *userService {
	return &userService{
		repository: r,
		config:     c,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repository.Users.Create(user)
}

func (s *userService) FindUserByEmail(email string) (*models.User, error) {
	return s.repository.Users.FindUserByEmail(email)
}

func (s *userService) FindAllUsers() ([]models.User, error) {
	return s.repository.Users.FindAll()
}

func (s *userService) FindUserByID(id uuid.UUID) (models.User, error) {
	return s.repository.Users.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repository.Users.Update(user)
}

func (s *userService) DeleteUserById(id uuid.UUID) error {
	return s.repository.Users.Delete(id)
}
