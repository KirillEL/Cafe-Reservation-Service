package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/types"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

type UpdateUserRequest struct {
	Login    string     `json:"login"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     types.Role `json:"role"`
}

func HandlerCommandsUpdateUserById(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("id")

		var req UpdateUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		userIDInt, _ := strconv.Atoi(userID)
		if err := service.UpdateUserById(uint(userIDInt), req.Login, req.Email, req.Password, req.Role); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "update user successfully"})
		return
	}
}
