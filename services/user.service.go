package services

import (
	"github.com/dduafa/go-server/core"
	"github.com/dduafa/go-server/models"
	"github.com/dduafa/go-server/repositories"
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

func (s *userService) CreateUser(payload *models.User) error {
	return s.repository.Users.Create(payload)

}

func (c *userService) FindUserByEmail(email string) (*models.User, error) {
	return c.repository.Users.FindUserByEmail(email)
}
