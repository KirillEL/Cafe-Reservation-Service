package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

func HandlerCommandsDeleteReservationById(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reservationID := ctx.Param("id")
		reservationIDINT, _ := strconv.Atoi(reservationID)
		isDeleted, err := service.DeleteReservationById(reservationIDINT)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !isDeleted {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "reservation ID not found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"reservation": "reservation was successfully deleted"})
		return

	}
}
