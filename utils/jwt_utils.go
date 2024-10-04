package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtKey)
}

// ValidateJWT verifies if the token is valid and not expired
func ValidateJWT(tokenStr string) bool {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        fmt.Println("Token parsing error:", err)
        return false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        exp := int64(claims["exp"].(float64))
        return exp >= time.Now().Unix()
    }

    return false
}
