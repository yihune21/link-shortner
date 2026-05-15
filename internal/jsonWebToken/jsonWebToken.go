package jsonwebtoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/config"
	"github.com/yihune21/link-shortner/internal/database"
)

func getJWTSecret() []byte {
	secret := config.GetEnv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return []byte(secret)
}

func GenerateAccessToken(user database.User) string {
	secretKey := getJWTSecret()

	claims := jwt.MapClaims{
		"sub":  user.ID.String(),
		"type": "access",
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}
    
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	return accessToken
}

func GenerateRefreshToken(user database.User) string {
	secretKey := getJWTSecret()

	claims := jwt.MapClaims{
		"sub":  user.ID.String(),
		"type": "refresh",
		"exp":  time.Now().Add(7 * 24 * 60).Unix(),
	}
    
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	return refreshToken
}

func GenerateToken(user database.User) string {
	return GenerateAccessToken(user)
}

func VerifyToken(tokenString string) bool {
	secretKey := getJWTSecret()
     
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Printf("Error with parsing: %v\n", err)
		return false
	}

	return parsedToken.Valid
}

func ExtractUserIDFromToken(tokenString string) (uuid.UUID, error) {
	secretKey := getJWTSecret()
    
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); ok {
			userID, err := uuid.Parse(sub)
			if err != nil {
				return uuid.Nil, err
			}
			return userID, nil
		}
		return uuid.Nil, fmt.Errorf("sub claim missing or invalid")
	}
	return uuid.Nil, fmt.Errorf("invalid token")
}

func VerifyRefreshToken(tokenString string) (uuid.UUID, error) {
	secretKey := getJWTSecret()
    
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
			return uuid.Nil, fmt.Errorf("invalid token type")
		}

		if sub, ok := claims["sub"].(string); ok {
			userID, err := uuid.Parse(sub)
			if err != nil {
				return uuid.Nil, err
			}
			return userID, nil
		}
		return uuid.Nil, fmt.Errorf("sub claim missing or invalid")
	}
	return uuid.Nil, fmt.Errorf("invalid token")
}