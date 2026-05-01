package main

import (
	"os"
	"scamshield-backend/app/handlers"
	"scamshield-backend/app/middleware"
	"scamshield-backend/app/services"
	"scamshield-backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize database
	db := config.InitDB()

	// Create router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize services
	urlAnalyzer := services.NewURLRiskAnalyzer()
	emailReviewer := services.NewEmailSecurityReviewer()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	scanHandler := handlers.NewScanHandler(db, urlAnalyzer)
	emailHandler := handlers.NewEmailReviewHandler(db, emailReviewer)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth routes
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Scan routes
		protected.POST("/scans", scanHandler.CreateScan)
		protected.GET("/scans", scanHandler.GetScans)
		protected.GET("/scans/:id", scanHandler.GetScan)

		// Email review routes
		protected.POST("/emails/review", emailHandler.ReviewEmail)
		protected.GET("/emails", emailHandler.GetReviews)
		protected.GET("/emails/:id", emailHandler.GetReview)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router.Run(":" + port)
}
