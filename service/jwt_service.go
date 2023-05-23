package service

import (
	"go-web-example/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type JwtService interface {
	CreateJwtToken(user model.User) (string, error)
	ValidateToken(tokenString string) (bool, error, string)
}

type JwtServiceImpl struct {
}

// We can create our custom claim , so we can add more data/struct type based on our need
type CustomClaims struct {
	// Our customClaims should embed type of claims
	StandardClaims jwt.StandardClaims
	User           model.User
}

func (j JwtServiceImpl) CreateJwtToken(user model.User) (string, error) {

	// 1. Claims = the data contained inside the JWT
	// Subject : to whom the token refer (you can give userId, username, sessionId , etc)
	// Expired : When Token expired
	// Issued At : When token issued
	// All this was standard as seen on jwt.io , you can add more or create your own
	standardClaims := jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	// 2. Secret Key
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	secretKey := os.Getenv("JWT_SECRET")

	// 3. token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)

	// 4. Create tge token string
	// Convert secret key to byte array
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JwtServiceImpl) ValidateToken(tokenString string) (bool, error, string) {

	// 1. Load secret key
	err := godotenv.Load()

	if err != nil {
		return false, err, ""
	}

	secretKey := os.Getenv("JWT_SECRET")

	// 2. Validate
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err, ""
	}

	// 3. Get Claims
	claims, ok := token.Claims.(*jwt.StandardClaims)

	if ok == false {
		return false, err, ""
	}

	/*
		54 Subject usually contain identifier to whom the token referred
		using sub , you could query the data in db , and then
		use the data to validate authorization (role, etc)
		you could do something like : userRepo.FindByUserName(subject)
	*/

	return true, nil, claims.Subject
}
