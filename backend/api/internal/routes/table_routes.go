package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes/table/handler/commands"
	"github.com/ivan/cafe_reservation/api/internal/routes/table/handler/queries"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
)

func NewTableRoutes(api *gin.RouterGroup, service table.UseCase) {
	api.POST("/create", commands.HandlerCommandsCreateTable(service))
	api.GET("/get/:id", queries.HandlerQueriesGetTableById(service))
	api.GET("/all", queries.HandlerQueriesGetAllTables(service))
	api.PUT("/update/:id", commands.HandlerCommandsUpdateTableById(service))
	api.DELETE("/delete/:id", commands.HandlerCommandsDeleteTableById(service))
}
