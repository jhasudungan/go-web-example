package model

import "net/http"

type FormData struct {
	Username string
	Email    string
	FullName string
}

type Product struct {
	ProductId    string  `json:"productId"`
	ProductName  string  `json:"productName"`
	ProductPrice float64 `json:"productPrice"`
}

type RestResponse struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func (r RestResponse) BuildSuccessResponse(payload interface{}) RestResponse {

	response := RestResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Payload: payload,
	}

	return response
}

type CreateProductRequest struct {
	NewProductName  string  `json:"newProductName"`
	NewProductPrice float64 `json:"newProductPrice"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
