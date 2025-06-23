package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return nil, err
	}

	hash := string(bytes)

	return &hash, nil
}

func ComparePassword(hashPW, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashPW), []byte(password))

	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateToken(email, id string) (*string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	secretKey := GoDotEnvVariable("JWT_SECRET")

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &token, nil
}

func VerifyToken(tokenString string) (*string, error) {
	secretKey := GoDotEnvVariable("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Error while validating token")
	}

	if !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Error while validating token")
	}
	userId, _ := claims["user_id"].(string)
	return &userId, nil
}

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("../../.env")

	if err != nil {

		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}
