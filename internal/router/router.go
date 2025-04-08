package router

import (
	"finance-tracker/internal/handler"
	"finance-tracker/internal/middleware"
	"finance-tracker/internal/repository"
	"finance-tracker/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewRouter sets up the routing for the API
func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Setup repositories
	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Setup usecases
	userUsecase := usecase.NewUserUsecase(userRepo)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)

	// Setup handlers
	userHandler := handler.NewUserHandler(userUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("/transactions", transactionHandler.Create)
		protected.GET("/transactions", transactionHandler.GetAll)
	}

	return r
}
