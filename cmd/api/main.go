package main

// @title Blockchain Certificate API
// @version 1.0
// @description API for registering and verifying blockchain-based certificates on Ethereum
// @description Allows registration of PDF certificates with metadata and verification through document hash

// @contact.name API Support
// @contact.email support@blockchain-cert.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/4nddrs/blockchain-cert/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/4nddrs/blockchain-cert/database"
	"github.com/4nddrs/blockchain-cert/internal/config"
	"github.com/4nddrs/blockchain-cert/internal/handlers"
	"github.com/4nddrs/blockchain-cert/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Connect to Ethereum client
	client, err := ethclient.Dial(cfg.AlchemyURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Initialize handlers
	handler := handlers.NewHandler(client, cfg)

	// Setup Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/openapi.json", "../../docs/swagger.json")

	// API v1 routes
	api := r.Group("/api/v1")
	{
		// Certificate endpoints
		api.POST("/register", handler.Register)
		api.POST("/verify", handler.Verify)
		api.GET("/certificates/:hash", handler.GetByHash)
		api.StaticFile("/openapi.json", "../../docs/swagger.json")
	}

	// Admin endpoints
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminAuth(cfg.AdminSecret))
	{
		admin.POST("/institutions", handler.CreateInstitution)
		admin.GET("/institutions", handler.ListInstitutions)
		admin.POST("/institutions/:id/credits", handler.AddCredits)
		admin.PUT("/institutions/:id/plan", handler.UpdatePlan)
	}

	// Start server
	serverAddr := ":" + cfg.ServerPort
	log.Printf("Server starting on http://localhost%s", serverAddr)
	log.Printf("Swagger documentation available at http://localhost%s/swagger/index.html", serverAddr)

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
