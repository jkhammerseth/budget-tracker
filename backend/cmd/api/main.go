package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/handler"
	"github.com/jkhammerseth/budget-tracker/backend/internal/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load("/Users/jonashammerseth/Desktop/småProsjekt/budget-tracker/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting the application...")

	router := gin.Default()

	// Setup CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes
	router.POST("/login", handler.Login)
	router.POST("/signup", handler.CreateUser)

	// Group private routes
	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		// Apply JWT middleware only to private routes
		api.GET("/users", handler.GetUsers)

		api.GET("/users/:id", handler.GetUser)
		api.PUT("/users/:id", handler.UpdateUser)
		api.DELETE("/users/:id", handler.DeleteUser)

		api.POST("/users/expense", handler.AddExpense)
		api.POST("/users/expenses", handler.AddExpenses)
		api.GET("/users/expenses", handler.GetExpenses)
		api.PUT("/users/expenses/:id", handler.UpdateExpense)
		api.DELETE("/users/expenses/:id", handler.DeleteExpense)

		api.POST("/users/incomes", handler.AddListIncome)
		api.POST("/users/income", handler.AddIncome)
		api.GET("/users/incomes", handler.GetIncomes)
		api.PUT("/users/incomes/:id", handler.UpdateIncome)
		api.DELETE("/users/incomes/:id", handler.DeleteIncome)

		api.POST("/users/loan", handler.AddLoan)
		api.GET("/users/loans", handler.GetLoan)
		api.PUT("/users/loans/:id", handler.UpdateLoan)
		api.DELETE("/users/loans/:id", handler.DeleteLoan)

		api.GET("/auth/status", handler.CheckAuthStatus)

		api.POST("/auth/logout", handler.Logout)

	}

	// Start the server
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
