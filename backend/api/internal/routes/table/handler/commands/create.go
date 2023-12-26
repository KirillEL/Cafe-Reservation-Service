package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

type CreateTableRequest struct {
	Number string `json:"number"`
	Seats  string `json:"seats"`
}

// TODO: /api/v1/table/create
func HandlerCommandsCreateTable(service table.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateTableRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		seatsInt, _ := strconv.Atoi(req.Seats)
		if err := service.CreateTable(req.Number, seatsInt); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "table was successfully created"})
		return
	}
}
