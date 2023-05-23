package service

import (
	"go-web-example/model"

	"github.com/google/uuid"
)

type ProductService interface {
	GetAllProduct() ([]model.Product, error)
	GetProductInformationById(productId string) (model.Product, error)
	CreateProduct(request model.CreateProductRequest) (model.Product, error)
}

type ProductServiceImpl struct {
	PointerToProductData *[]model.Product
}

func (p ProductServiceImpl) GetAllProduct() ([]model.Product, error) {

	// Get the real value from pointer
	data := *p.PointerToProductData

	return data, nil
}

func (p ProductServiceImpl) GetProductInformationById(productId string) (model.Product, error) {

	result := model.Product{}

	// Get the real value from pointer
	data := *p.PointerToProductData

	for _, product := range data {

		if product.ProductId == productId {
			result = product
		}

	}

	return result, nil

}

func (p ProductServiceImpl) CreateProduct(request model.CreateProductRequest) (model.Product, error) {

	// Create new struct
	product := model.Product{
		ProductId:    uuid.NewString(),
		ProductName:  request.NewProductName,
		ProductPrice: request.NewProductPrice,
	}

	// Access the underlying data through pointer
	data := *p.PointerToProductData
	data = append(data, product)

	return product, nil
}
