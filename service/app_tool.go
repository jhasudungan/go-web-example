package service

import (
	"encoding/json"
	"go-web-example/model"
	"io/ioutil"
)

type ApplicationTool interface {
	LoadAllInitialProductData(filePath string) []model.Product
	LoadAllInitialUserData(filePath string) []model.User
}

type ApplicationToolImpl struct{}

func (a ApplicationToolImpl) LoadAllInitialProductData(filePath string) []model.Product {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic("Error Load Initial Product Data : " + err.Error())
	}

	result := make([]model.Product, 0)

	// unmarshal json
	err = json.Unmarshal(data, &result)

	return result
}

func (a ApplicationToolImpl) LoadAllInitialUserData(filePath string) []model.User {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic("Error Load Initial User Data : " + err.Error())
	}

	result := make([]model.User, 0)

	// unmarshal json
	err = json.Unmarshal(data, &result)

	return result
}
