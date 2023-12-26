package commands

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/middlewares"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func HandlerCommandLoginUser(service user.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := service.Login(req.Email, req.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := middlewares.CreateToken(user.ID, user.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.SetCookie("token", token, 3600, "/", "", false, false)

		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"token":  token,
		})
		return
	}

}
