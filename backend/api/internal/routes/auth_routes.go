package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes/auth/handler/commands"
	"github.com/ivan/cafe_reservation/api/internal/routes/auth/handler/queries"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

func NewAuthRoutes(api *gin.RouterGroup, service user.UseCase) {
	api.POST("/register", commands.HandlerCommandRegisterUser(service))
	api.GET("/verify", queries.HandlerQueriesVerifyUser(service))
	api.POST("/login", commands.HandlerCommandLoginUser(service))
	api.DELETE("/logout", func(ctx *gin.Context) {
		ctx.SetCookie("token", "", -1, "/", ".", false, false)
		ctx.JSON(http.StatusOK, gin.H{"message": "success logout"})
		return
	})
}
