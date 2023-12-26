package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

type UpdateTableRequest struct {
	Number string `json:"number"`
	Seats  int    `json:"seats"`
}

func HandlerCommandsUpdateTableById(service table.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableID := ctx.Param("id")
		var req UpdateTableRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		tableIDInt, _ := strconv.Atoi(tableID)
		if err := service.UpdateTableById(uint64(tableIDInt), req.Number, req.Seats); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Updated table successfully"})
		return
	}
}
