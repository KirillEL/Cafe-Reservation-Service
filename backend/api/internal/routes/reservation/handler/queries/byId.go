package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

func HandlerQueriesGetReservationsByUserId(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userID")
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found"})
			return
		}
		reservations, err := service.GetReservationsByUserId(userID.(uint))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": reservations})
		return
	}
}
