package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

func HandlerCommandsDeleteUserById(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("id")
		userIDInt, _ := strconv.Atoi(userID)
		if err := service.DeleteUserById(uint(userIDInt)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User was successfully deleted"})
		return
	}
}
