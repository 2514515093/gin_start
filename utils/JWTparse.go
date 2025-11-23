package utils

import (
	"fmt"
	"gin_start/repository"

	"github.com/dgrijalva/jwt-go"
)

func ParseTokenToUser(tokenString string) (*repository.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("1234"), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &repository.User{}
		// 注意 jwt 的数字默认是 float64，需要转换
		if idFloat, ok := claims["id"].(float64); ok {
			user.ID = uint(idFloat)
		}
		if username, ok := claims["username"].(string); ok {
			user.Username = username
		}
		return user, nil
	}
	return nil, fmt.Errorf("invalid token")
}
