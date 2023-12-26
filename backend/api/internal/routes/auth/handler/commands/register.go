package commands

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

type RegisterUserRequest struct {
	Login    string `json:"login" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func HandlerCommandRegisterUser(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RegisterUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		login := req.Login
		email := req.Email
		password := req.Password
		_, err := service.Register(login, email, password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
		return

	}
}
