package service

import "go-web-example/model"

type DataService interface {
	FindByUsername(username string) model.User
}

type DataServiceImpl struct {
	UserData []model.User
}

func (d DataServiceImpl) FindByUsername(username string) model.User {

	var result model.User

	for _, user := range d.UserData {

		if user.Username == username {
			result = user
		}
	}

	return result
}
