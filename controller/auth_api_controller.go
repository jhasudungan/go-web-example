package controller

import (
	"encoding/json"
	"go-web-example/model"
	"go-web-example/service"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthApiController struct {
	AuthService service.AuthService
	JwtService  service.JwtService
	DataService service.DataService
}

/*
	Login API
*/
func (a AuthApiController) HanldeLoginAPI(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	loginRequest := model.LoginRequest{}

	// Unmarshall
	err = json.Unmarshal(body, &loginRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err, user := a.AuthService.Authenticate(loginRequest)

	if result == false {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JWT Token Creation
	token, err := a.JwtService.CreateJwtToken(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response data
	responseData := make(map[string]string)
	responseData["token"] = token
	responseData["message"] = "success"

	response := model.RestResponse{}.BuildSuccessResponse(responseData)
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

func (a AuthApiController) GetDataFromToken(w http.ResponseWriter, r *http.Request) {

	// Read token from "Bearer "
	authorization := r.Header.Get("Authorization")

	// Split into array : "[Bearer] [abcdxeyzhij1234]"
	tokenParts := strings.Split(authorization, " ")

	// Validate to format
	if len(tokenParts) != 2 {
		http.Error(w, "invalid token format", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(tokenParts[0]) != "bearer" {
		http.Error(w, "invalid token format", http.StatusInternalServerError)
		return
	}

	// Get the token
	token := tokenParts[1]

	valid, err, subject := a.JwtService.ValidateToken(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if valid == false {
		http.Error(w, "token is not valid", http.StatusInternalServerError)
		return
	}

	// using the subject (username) , we find the real data
	user := a.DataService.FindByUsername(subject)

	response := model.RestResponse{}.BuildSuccessResponse(user)
	jsonResponse, err := json.Marshal(response)

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
