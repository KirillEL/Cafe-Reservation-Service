package commands

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
)

type UpdateCustomerRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Phone  int    `json:"phone"`
	Email  string `json:"email"`
}

func HandlerCommandsUpdateCustomerById(service customer.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerID := ctx.Param("id")

		var req UpdateCustomerRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		customerIDInt, _ := strconv.Atoi(customerID)
		if err := service.UpdateCustomerById(uint(customerIDInt), uint(req.UserID), req.Name, req.Phone, req.Email); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
		return
	}
}
