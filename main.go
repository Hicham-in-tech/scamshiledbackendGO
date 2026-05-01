package main

import (
	"net/http"
	"os"
	"sync"
	"scamshield-backend/app/handlers"
	"scamshield-backend/app/middleware"
	"scamshield-backend/app/services"
	"scamshield-backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router     *gin.Engine
	routerOnce sync.Once
)

func getRouter() *gin.Engine {
	routerOnce.Do(func() {
		// Load environment variables
		godotenv.Load()

		// Initialize database
		db := config.InitDB()

		// Create router
		r := gin.Default()

		// CORS configuration
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://*.vercel.app"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization", "X-User-ID"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowWildcard:    true,
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
		r.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// Auth routes
		r.POST("/auth/register", authHandler.Register)
		r.POST("/auth/login", authHandler.Login)

		// Protected routes
		protected := r.Group("/api")
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

		router = r
	})

	return router
}

// Handler is the exported entrypoint Vercel expects for Go functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	getRouter().ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, getRouter())
}
