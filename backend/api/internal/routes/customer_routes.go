package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes/customer/handler/commands"
	"github.com/ivan/cafe_reservation/api/internal/routes/customer/handler/queries"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
)

func NewCustomerRoutes(api *gin.RouterGroup, service customer.UseCase) {
	api.POST("/create", commands.HandlerCommandsCreateCustomer(service))
	api.DELETE("/delete/:id", commands.HandlerCommandsDeleteCustomerById(service))
	api.PUT("/update/:id", commands.HandlerCommandsUpdateCustomerById(service))
	api.GET("/all", queries.HandlerQueriesGetAllCustomers(service))
}
