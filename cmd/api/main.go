package main

// export GOTMPDIR=~/go-cache/tmp

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
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/4nddrs/blockchain-cert/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
)

var client *ethclient.Client

func main() {
	// 1. Load environment variables
	godotenv.Load("../../.env")

	var err error
	client, err = ethclient.Dial(os.Getenv("ALCHEMY_URL"))

	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// 2. Config Server
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/openapi.json", "../../docs/swagger.json")
	// CORS Config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", handleRegister)
		v1.POST("/verify", handleVerify)
		v1.GET("/certificates/:hash", handleGetByHash)
		v1.StaticFile("/openapi.json", "../../docs/swagger.json")
	}

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}

// handleRegister godoc
// @Summary Register a new certificate on the blockchain
// @Description Uploads a PDF certificate, generates its SHA256 hash, and registers it on Ethereum with student metadata
// @Description The certificate is permanently stored on the blockchain and can be verified later
// @Tags certificates
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF certificate file to register"
// @Param student_name formData string true "Full name of the student receiving the certificate"
// @Param course_name formData string true "Name of the course or program completed"
// @Param issuer formData string true "Name of the institution or organization issuing the certificate"
// @Success 200 {object} map[string]interface{} "Certificate registered successfully with transaction hash"
// @Failure 400 {object} map[string]interface{} "Failed to register certificate on blockchain"
// @Failure 500 {object} map[string]interface{} "Internal server error (file save failed)"
// @Router /register [post]
func handleRegister(c *gin.Context) {
	// Input data structure
	file, _ := c.FormFile("pdf")
	studentName := c.PostForm("student_name")
	courseName := c.PostForm("course_name")
	issuer := c.PostForm("issuer")

	// Save temporary file to hash
	tempFilePath := "./temp_/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	defer os.Remove(tempFilePath)

	// Blockchain interaction

	hash, _ := blockchain.GenerateHash(tempFilePath)
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	privateKey := os.Getenv("PRIVATE_KEY")

	txHash, err := blockchain.RegisterCertificate(client, privateKey, contractAddress, hash, studentName, courseName, issuer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register certificate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"hash":    hash,
		"tx_hash": txHash,
	})
}

// handleVerify godoc
// @Summary Verify a certificate by uploading the PDF
// @Description Uploads a PDF file, generates its hash, and checks if it exists on the blockchain
// @Description Returns certificate details if found and valid
// @Tags certificates
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF certificate file to verify"
// @Success 200 {object} map[string]interface{} "Certificate is valid with full details (student_name, course_name, issuer, date)"
// @Failure 404 {object} map[string]interface{} "Certificate not found or invalid"
// @Router /verify [post]
func handleVerify(c *gin.Context) {
	file, _ := c.FormFile("pdf")

	tempPath := "./verify_" + file.Filename
	c.SaveUploadedFile(file, tempPath)
	defer os.Remove(tempPath)

	hash, _ := blockchain.GenerateHash(tempPath)
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	cert, err := blockchain.VerifyCertificate(client, contractAddress, hash)

	if err != nil || !cert.IsValid {
		c.JSON(http.StatusNotFound, gin.H{"status": "not_found", "message": "Certificate not found or invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "valid",
		"student_name": cert.StudentName,
		"course_name":  cert.CourseName,
		"issuer":       cert.IssuerName,
		"date":         cert.DateEmited.Format("2006-01-02 15:04:05"),
	})
}

// handleGetByHash godoc
// @Summary Get certificate details by hash
// @Description Retrieves certificate information from the blockchain using its SHA256 hash
// @Description No file upload required - only the hash is needed for lookup
// @Tags certificates
// @Accept json
// @Produce json
// @Param hash path string true "SHA256 hash of the certificate (66 characters including 0x prefix)" minlength(66) maxlength(66)
// @Success 200 {object} map[string]interface{} "Certificate found with full details"
// @Failure 400 {object} map[string]interface{} "Invalid hash format (must be 66 characters)"
// @Failure 404 {object} map[string]interface{} "Certificate not found or invalid"
// @Failure 500 {object} map[string]interface{} "Failed to retrieve certificate from blockchain"
// @Router /certificates/{hash} [get]
func handleGetByHash(c *gin.Context) {

	// Get hash from URL
	hash := c.Param("hash")

	//Validate hash format
	if len(hash) != 66 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid hash format"})
		return
	}

	// Consume blockchain
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	cert, err := blockchain.VerifyCertificate(client, contractAddress, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to retrieve certificate"})
		return
	}

	// Verify Certificate
	if !cert.IsValid {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "not_found",
			"message": "Certificate not found or invalid",
		})
		return
	}

	// Return certificate details
	c.JSON(http.StatusOK, gin.H{
		"status": "valid",
		"data": gin.H{
			"hash":         hash,
			"student_name": cert.StudentName,
			"course_name":  cert.CourseName,
			"issuer":       cert.IssuerName,
			"date":         cert.DateEmited.Format("2006-01-02 15:04:05"),
		},
	})
}
