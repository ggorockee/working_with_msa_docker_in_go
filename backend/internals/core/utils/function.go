package utils

import "github.com/golang-jwt/jwt/v5"

func ValidToken(t *jwt.Token, id int) bool {

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == id
}

func GetUserIdFromJwtToken(token *jwt.Token) int {
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))
	return userId
}