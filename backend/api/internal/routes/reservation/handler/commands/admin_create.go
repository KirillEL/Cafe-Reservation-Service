package commands

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

type ReservationCreateRequest struct {
	CustomerID  int    `json:"customer_id"`
	TableID     int    `json:"table_id"`
	ReserveTime string `json:"reserve_time"`
	Duration    int    `json:"duration"`
	Status      string `json:"status"`
}

func HandlerCommandsCreateReservationByAdmin(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ReservationCreateRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := service.CreateReservationByAdmin(uint(req.CustomerID), uint(req.TableID), req.ReserveTime, req.Duration, req.Status); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully"})
		return

	}
}
