package commands

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

type CreateReservationRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	TableNumber string `json:"table_number"`
	ReserveTime string `json:"reserve_time"`
	Duration    string `json:"duration"`
}

func HandlerCommandsCreateReservation(service reservation.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userID")
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found"})
			return
		}
		log.Println("here")

		var req CreateReservationRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(req)
		log.Println(req.Duration)
		err := service.CreateReservation(userID.(uint), req.Name, req.Phone, req.Email, req.TableNumber, req.ReserveTime, req.Duration)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully"})
		return
	}
}
