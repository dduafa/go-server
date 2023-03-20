package services

import (
	"github.com/dduafa/go-server/core"
	"github.com/dduafa/go-server/repositories"
)

type Services struct {
	Users *userService
}

func NewService(r repositories.Repo, c *core.Config) Services {
	return Services{
		Users: newUserService(r, c),
	}
}
