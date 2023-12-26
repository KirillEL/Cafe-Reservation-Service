package commands

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/types"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

type CreateUserRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func HandlerCommandsCreateUser(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Role := types.Role(req.Role)
		if err := service.CreateUser(req.Login, req.Email, req.Password, Role); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "successfully created user"})
		return
	}
}
