package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
)

func HandlerQueriesGetAllCustomers(service customer.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := service.GetAllCustomers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": customers})
		return
	}
}
