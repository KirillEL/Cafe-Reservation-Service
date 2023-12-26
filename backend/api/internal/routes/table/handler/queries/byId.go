package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

func HandlerQueriesGetTableById(service table.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
			return
		}
		table, err := service.GetTableById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": table})
		return
	}
}
