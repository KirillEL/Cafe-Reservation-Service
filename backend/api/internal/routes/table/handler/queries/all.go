package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

func HandlerQueriesGetAllTables(service table.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tables, err := service.GetAvailableTables()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": tables})
		return
	}
}
