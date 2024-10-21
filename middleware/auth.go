package middleware

import (
	"net/http"

	"github.com/Izzat-Khudoyberganov/dictionary-app/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Next()
}
