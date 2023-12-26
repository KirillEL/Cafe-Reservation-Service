package commands

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

type UpdateReservationRequest struct {
	CustomerID  int       `json:"customer_id"`
	TableID     int       `json:"table_id"`
	ReserveTime time.Time `json:"reserve_time"`
	Duration    int       `json:"duration"`
	Status      string    `json:"status"`
}

func HandlerCommandsUpdateReservationByCustomerId(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reservationID := ctx.Param("id")

		var req UpdateReservationRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		reservationIDInt, _ := strconv.Atoi(reservationID)
		if err := service.UpdateReservationById(uint(reservationIDInt), uint(req.CustomerID), uint(req.TableID), req.ReserveTime, req.Duration, req.Status); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Reservation updated successfully"})
		return

	}
}
