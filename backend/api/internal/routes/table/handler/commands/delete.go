package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

func HandlerCommandsDeleteTableById(service table.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableID := ctx.Param("id")
		tableIDInt, _ := strconv.Atoi(tableID)
		if err := service.DeleteTableById(uint(tableIDInt)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "delete table successfully"})
		return
	}
}
