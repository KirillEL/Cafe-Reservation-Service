package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
)

func HandlerCommandsDeleteCustomerById(service customer.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerID := ctx.Param("id")
		customerIDInt, _ := strconv.Atoi(customerID)
		if err := service.DeleteCustomerById(uint(customerIDInt)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
		return
	}
}
