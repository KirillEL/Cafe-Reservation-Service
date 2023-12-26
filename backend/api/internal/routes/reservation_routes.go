package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes/reservation/handler/commands"
	"github.com/ivan/cafe_reservation/api/internal/routes/reservation/handler/queries"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
)

func NewReservationRoutes(api *gin.RouterGroup, service reservation.UseCase) {
	api.POST("/create", commands.HandlerCommandsCreateReservation(service))
	api.POST("/admin/create", commands.HandlerCommandsCreateReservationByAdmin(service))
	api.PUT("/update/:id", commands.HandlerCommandsUpdateReservationByCustomerId(service))
	api.DELETE("/delete/:id", commands.HandlerCommandsDeleteReservationById(service))
	api.GET("/get/own", queries.HandlerQueriesGetReservationsByUserId(service))
	api.GET("/all", queries.HandlerQueriesGetAllReservations(service))
}
