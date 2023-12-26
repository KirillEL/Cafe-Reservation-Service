package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ivan/cafe_reservation/internal/types"
	"github.com/ivan/cafe_reservation/pkg/config"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "JWT token is required in the cookie"})
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error reading token from cookie"})
			}
			return
		}

		isValid, claims := ValidateJWTToken(jwtToken)
		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}

type CustomClaims struct {
	UserID uint       `json:"user_id"`
	Role   types.Role `json:"role"`
	jwt.StandardClaims
}

type ValidateClaimsReturn struct {
	UserID uint       `json:"user_id"`
	Role   types.Role `json:"role"`
}

func ValidateJWTToken(tokenString string) (bool, ValidateClaimsReturn) {
	var jwtSecretKey = []byte(config.Env.JwtSecret)
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	t := ValidateClaimsReturn{
		UserID: claims.UserID,
		Role:   claims.Role,
	}

	if err != nil {
		return false, ValidateClaimsReturn{}
	}

	if !token.Valid {
		return false, ValidateClaimsReturn{}
	}

	return true, t
}

type Claims struct {
	UserID uint       `json:"user_id"`
	Role   types.Role `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(userID uint, role types.Role) (string, error) {
	var jwtSecretKey = []byte(config.Env.JwtSecret)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
