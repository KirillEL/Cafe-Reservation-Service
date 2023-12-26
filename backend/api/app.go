package api

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ivan/cafe_reservation/api/internal/routes"
	"github.com/ivan/cafe_reservation/api/middlewares"
	"github.com/ivan/cafe_reservation/internal/database"
	"github.com/ivan/cafe_reservation/internal/repository"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
	"github.com/ivan/cafe_reservation/pkg/config"
	"gorm.io/driver/postgres"
)

func Run(port int) {
	config.LoadEnvFile("cmd/.env")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DBHost, config.Env.DBPort, config.Env.DBUser, config.Env.DBPassword, config.Env.DBName)
	gormDB, err := database.ConnectToDB(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	app.Use(gin.Logger())
	app.Use(gin.ErrorLogger())
	app.Use(cors.New(cors.Config{
		// Allow all origins for development purposes
		// In production, replace with your frontend's actual URL
		AllowOrigins: []string{"http://localhost:3000"},

		// Allow only specific methods
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		// Allow specific headers
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},

		// Indicate whether the response can include credentials like cookies or HTTP auth
		AllowCredentials: true,

		// Pre-flight request cache duration
		MaxAge: 12 * time.Hour,
	}))
	//app.Use(compress.New())

	userRepository := repository.NewUserRepository(gormDB)
	reservationRepository := repository.NewReservationRepository(gormDB)
	tableRepository := repository.NewTableRepository(gormDB)
	customerRepository := repository.NewCustomerRepository(gormDB)

	// services
	userService := user.NewUserService(userRepository)
	reservationService := reservation.NewReservationService(reservationRepository)
	tableService := table.NewTableService(tableRepository)
	customerService := customer.NewCustomerService(customerRepository)

	api := app.Group("/api")
	api_v1 := api.Group("/v1")

	api_v1.Use(middlewares.JWTAuthMiddleware())

	routes.NewAuthRoutes(api.Group("/auth"), userService)
	routes.NewTableRoutes(api_v1.Group("/table"), tableService)
	routes.NewCustomerRoutes(api_v1.Group("/customer"), customerService)
	routes.NewReservationRoutes(api_v1.Group("/reservation"), reservationService)
	routes.NewUserRoutes(api_v1.Group("/user"), userService)

	// app.Any("*", func(c *gin.Context) {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"message": "Not Found",
	// 	})
	// })

	app.Run("0.0.0.0:1010")

}
