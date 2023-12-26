package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

func HandlerQueriesGetAllReservations(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reservations, err := service.GetAllReservations()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": reservations})
		return
	}
}
