package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

func HandlerQueriesGetAllUsers(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := service.GetAllUsers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": users})
		return
	}
}
