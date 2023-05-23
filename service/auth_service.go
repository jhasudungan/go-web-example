package service

import (
	"errors"
	"go-web-example/model"
)

type AuthService interface {
	Authenticate(request model.LoginRequest) (bool, error, model.User)
}

type AuthServiceImpl struct {
	DataService DataService
}

func (a AuthServiceImpl) Authenticate(request model.LoginRequest) (bool, error, model.User) {

	// 1. Find User by Username
	attemptedUser := a.DataService.FindByUsername(request.Username)

	if attemptedUser == (model.User{}) {
		return false, errors.New("authenticate failed - user not found"), model.User{}
	}

	// 2. Compare Password
	if attemptedUser.Password != request.Password {
		return false, errors.New("authenticate failed - wrong password"), model.User{}
	}

	return true, nil, attemptedUser
}
