package main

import (
	"go-web-example/controller"
	"go-web-example/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	// Initial Cookie Data
	err := godotenv.Load()

	if err != nil {
		panic("Please check cookie store : " + err.Error())
	}

	key := []byte(os.Getenv("GORILLA_SESSION_SECRET"))
	pointerToSessionStore := sessions.NewCookieStore(key)

	// Load Initial Data
	applicationTool := service.ApplicationToolImpl{}

	initialProductData := applicationTool.LoadAllInitialProductData("data-product.json")
	initialUserData := applicationTool.LoadAllInitialUserData("data-user.json")

	// Service
	productService := service.ProductServiceImpl{
		PointerToProductData: &initialProductData,
	}

	dataService := service.DataServiceImpl{
		UserData: initialUserData,
	}

	authService := service.AuthServiceImpl{
		DataService: dataService,
	}

	jwtService := service.JwtServiceImpl{}

	// Controller
	publicController := controller.PublicController{}

	productController := controller.ProductController{
		ProductService: productService,
	}

	authController := controller.AuthController{
		AuthService:           authService,
		JwtService:            jwtService,
		DataService:           dataService,
		PointerToSessionStore: pointerToSessionStore}

	authApiController := controller.AuthApiController{
		AuthService: authService,
		JwtService:  jwtService,
		DataService: dataService,
	}

	// WEB-Route
	router.HandleFunc("/test-page", publicController.ShowTestPage).Methods("GET")
	router.HandleFunc("/", publicController.ShowIndexPage).Methods("GET")
	router.HandleFunc("/form", publicController.ShowFormPage).Methods("GET")
	router.HandleFunc("/form-post", publicController.HandleFormSubmit).Methods("POST")
	router.HandleFunc("/about", publicController.ShowAboutPage).Methods("GET")

	// REST API
	router.HandleFunc("/api/v1/product", productController.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/v1/product/{productId}", productController.GetProductById).Methods("GET")
	router.HandleFunc("/api/v1/product", productController.CreateProduct).Methods("POST")

	// Auth and Session
	router.HandleFunc("/login", authController.ShowLoginForm).Methods("GET")
	router.HandleFunc("/login/submit", authController.HandleLoginSubmit).Methods("POST")
	router.HandleFunc("/login/success", authController.ShowLoginSuccess).Methods("GET")

	// Login API and JWT
	router.HandleFunc("/api/v1/login", authApiController.HanldeLoginAPI).Methods("POST")
	router.HandleFunc("/api/v1/me", authApiController.GetDataFromToken).Methods("GET")

	log.Print("Web run on server using port :8081 ")

	err = http.ListenAndServe(":8081", router)

	if err != nil {
		log.Printf(" Error : %v ", err.Error())
	}

}
