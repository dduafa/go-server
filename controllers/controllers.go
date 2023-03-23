package controllers

import (
	"github.com/dduafa/go-server/services"
)

type Controllers struct {
	Auth *authController
	User *userController
}

func NewController(s services.Services) Controllers {
	return Controllers{
		Auth: newAuthController(s),
		User: newUserController(s),
	}
}
