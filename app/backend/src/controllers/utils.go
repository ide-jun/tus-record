package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func checkAuthHeader(context *gin.Context) (*jwt.Token, bool) {
	// ヘッダーからトークンを取得
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		return nil, false
	}

	// Bearer スキーマをチェック
	splitted := strings.Split(authHeader, " ")
	if len(splitted) != 2 || splitted[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Authorization header format",
		})
		return nil, false
	}

	tokenString := splitted[1]

	// トークンの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return nil, false
	}

	return token, true
}
