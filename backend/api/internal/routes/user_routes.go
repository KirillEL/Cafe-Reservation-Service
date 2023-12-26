package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes/user/handler/commands"
	"github.com/ivan/cafe_reservation/api/internal/routes/user/handler/queries"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

func NewUserRoutes(api *gin.RouterGroup, service user.UseCase) {
	api.GET("/all", queries.HandlerQueriesGetAllUsers(service))
	api.POST("/create", commands.HandlerCommandsCreateUser(service))
	api.PUT("/update/:id", commands.HandlerCommandsUpdateUserById(service))
	api.DELETE("/delete/:id", commands.HandlerCommandsDeleteUserById(service))
}
