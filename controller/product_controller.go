package controller

import (
	"encoding/json"
	"go-web-example/model"
	"go-web-example/service"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductController struct {
	ProductService service.ProductService
}

/*
Send JSON back
*/
func (p ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {

	// 1. Get data from service
	data, err := p.ProductService.GetAllProduct()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Create json response using marshall
	response := model.RestResponse{}.BuildSuccessResponse(data)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

/*
Getting the path variable
*/
func (p ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {

	// 1. Get path variable
	vars := mux.Vars(r)
	productId := vars["productId"]

	// 2. Call service
	data, err := p.ProductService.GetProductInformationById(productId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Create json response
	response := model.RestResponse{}.BuildSuccessResponse(data)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

/*
Handling the response body
*/
func (p ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	// 1. Read the JSON Body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	// 2. Create Struct
	createProductRequest := model.CreateProductRequest{}

	// 3. Unmarshall json -> struct
	err = json.Unmarshal(body, &createProductRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Call the service , do the bussiness logic
	newProduct, err := p.ProductService.CreateProduct(createProductRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. Create json response
	response := model.RestResponse{}.BuildSuccessResponse(newProduct)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
