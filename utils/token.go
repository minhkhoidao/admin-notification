package utils

import (
	"admin-backend/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &models.JwtCustomClaims{
		Name: user.Username,
		ID:   strconv.Itoa(user.ID), // Convert integer ID to string
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claimsRefresh := &models.JwtCustomRefreshClaims{
		ID: strconv.Itoa(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	log.Println("Token: ", requestToken)
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	log.Println("Claims:", claims)

	id, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("user not found")
	}

	return id, nil
}
