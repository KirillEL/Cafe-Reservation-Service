package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/middlewares"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

func HandlerQueriesVerifyUser(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error reading token"})
			}
			return
		}

		// Validate the token
		isValid, claims := middlewares.ValidateJWTToken(token)
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Token is valid", "claims": claims})
		return
	}
}
