package main

// export GOTMPDIR=~/go-cache/tmp

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

	}

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}

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
