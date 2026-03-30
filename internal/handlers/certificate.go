package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
	"github.com/4nddrs/blockchain-cert/internal/config"
	"github.com/4nddrs/blockchain-cert/internal/repository"
	"github.com/4nddrs/blockchain-cert/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	Client *ethclient.Client
	Config *config.Config
}

// NewHandler creates a new Handler instance
func NewHandler(client *ethclient.Client, cfg *config.Config) *Handler {
	return &Handler{
		Client: client,
		Config: cfg,
	}
}

// RegisterResponse represents the response for certificate registration
type RegisterResponse struct {
	Status string `json:"status" example:"success"`
	Hash   string `json:"hash" example:"0xabcd1234..."`
	TxHash string `json:"tx_hash" example:"0x1234abcd..."`
}

// VerifyResponse represents the response for certificate verification
type VerifyResponse struct {
	Status      string `json:"status" example:"valid"`
	StudentName string `json:"student_name" example:"John Doe"`
	CourseName  string `json:"course_name" example:"Blockchain Development"`
	Issuer      string `json:"issuer" example:"Tech University"`
	Date        string `json:"date" example:"2026-03-27 20:00:00"`
}

// CertificateDetailsResponse represents detailed certificate information
type CertificateDetailsResponse struct {
	Status string               `json:"status" example:"valid"`
	Data   CertificateData      `json:"data"`
}

// CertificateData holds certificate details
type CertificateData struct {
	Hash        string `json:"hash" example:"0xabcd1234..."`
	StudentName string `json:"student_name" example:"John Doe"`
	CourseName  string `json:"course_name" example:"Blockchain Development"`
	Issuer      string `json:"issuer" example:"Tech University"`
	Date        string `json:"date" example:"2026-03-27 20:00:00"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Certificate not found"`
	Message string `json:"message,omitempty" example:"Additional error details"`
}

// Register godoc
// @Summary Register a new certificate on the blockchain
// @Description Uploads a PDF certificate, generates its SHA256 hash, and registers it on Ethereum with student metadata. The certificate is permanently stored on the blockchain and can be verified later.
// @Tags certificates
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF certificate file to register"
// @Param student_name formData string true "Full name of the student receiving the certificate"
// @Param course_name formData string true "Name of the course or program completed"
// @Param issuer formData string true "Name of the institution or organization issuing the certificate"
// @Param fingerprint formData string false "Browser fingerprint for trial usage tracking"
// @Param X-API-Key header string false "Institution API key (if not provided, trial mode is used)"
// @Success 200 {object} RegisterResponse "Certificate registered successfully"
// @Failure 400 {object} ErrorResponse "Invalid request or registration failed"
// @Failure 401 {object} ErrorResponse "Invalid API key"
// @Failure 402 {object} ErrorResponse "Insufficient credits"
// @Failure 409 {object} ErrorResponse "Certificate already registered"
// @Failure 429 {object} ErrorResponse "Free trial limit reached"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	apiKey := c.GetHeader("X-API-Key")
	fingerprint := c.PostForm("fingerprint")
	clientIP := c.ClientIP()

	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PDF file is required"})
		return
	}

	studentName := c.PostForm("student_name")
	courseName := c.PostForm("course_name")
	issuer := c.PostForm("issuer")

	// Validate required fields
	if studentName == "" || courseName == "" || issuer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_name, course_name, and issuer are required"})
		return
	}

	// Determine if trial or institution-based
	var institutionID string
	isTrial := false

	if apiKey == "" {
		// Trial mode
		eligible, err := repository.CheckTrialEligibility(c.Request.Context(), clientIP, fingerprint)
		if err != nil {
			log.Printf("Error checking trial eligibility: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate trial eligibility"})
			return
		}
		if !eligible {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Free trial limit reached"})
			return
		}
		isTrial = true
	} else {
		// Institution mode
		inst, err := repository.GetInstitutionByAPIKey(c.Request.Context(), apiKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}
		if inst.CreditsRemaining <= 0 {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Insufficient credits. Please contact support."})
			return
		}
		institutionID = inst.ID.String()
	}

	// Save file temporarily
	tempFilePath := h.Config.TempUploadDir + "/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}
	defer os.Remove(tempFilePath)

	// Generate hash
	hash, err := blockchain.GenerateHash(tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file hash"})
		return
	}

	// Check if certificate already exists
	existingCert, err := repository.GetCertificateByHash(c.Request.Context(), hash)
	if err != nil {
		log.Printf("Error checking existing certificate: %v", err)
	}
	if existingCert != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":   "Certificate already registered",
			"tx_hash": existingCert.TxHash,
		})
		return
	}

	// Register on blockchain
	txHash, err := blockchain.RegisterCertificate(
		h.Client,
		h.Config.PrivateKey,
		h.Config.ContractAddress,
		hash,
		studentName,
		courseName,
		issuer,
	)
	if err != nil {
		log.Printf("Blockchain registration failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register certificate on blockchain"})
		return
	}

	// Save to database
	certData := models.Certificate{
		FileHash:    hash,
		StudentName: studentName,
		CourseName:  courseName,
		TxHash:      txHash,
	}

	if err := repository.SaveCertificateAndMarkUsage(c.Request.Context(), certData, clientIP, fingerprint, institutionID, isTrial); err != nil {
		log.Printf("Error saving certificate to database: %v", err)
		// Note: Certificate is already on blockchain, so we return success
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"hash":    hash,
		"tx_hash": txHash,
	})
}

// Verify godoc
// @Summary Verify a certificate by uploading the PDF
// @Description Uploads a PDF file, generates its hash, and checks if it exists on the blockchain. Returns certificate details if found and valid.
// @Tags certificates
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF certificate file to verify"
// @Success 200 {object} VerifyResponse "Certificate is valid"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 404 {object} ErrorResponse "Certificate not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /verify [post]
func (h *Handler) Verify(c *gin.Context) {
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PDF file is required"})
		return
	}

	tempPath := h.Config.TempUploadDir + "/verify_" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}
	defer os.Remove(tempPath)

	hash, err := blockchain.GenerateHash(tempPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file hash"})
		return
	}

	cert, err := blockchain.VerifyCertificate(h.Client, h.Config.ContractAddress, hash)
	if err != nil {
		log.Printf("Error verifying certificate: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify certificate"})
		return
	}

	if !cert.IsValid {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "not_found",
			"message": "Certificate not found or invalid",
		})
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

// GetByHash godoc
// @Summary Get certificate details by hash
// @Description Retrieves certificate information from the blockchain using its SHA256 hash. No file upload required.
// @Tags certificates
// @Accept json
// @Produce json
// @Param hash path string true "SHA256 hash of the certificate (must be 66 characters: 0x + 64 hex digits)" minlength(66) maxlength(66)
// @Success 200 {object} CertificateDetailsResponse "Certificate found"
// @Failure 400 {object} ErrorResponse "Invalid hash format"
// @Failure 404 {object} ErrorResponse "Certificate not found"
// @Failure 500 {object} ErrorResponse "Failed to retrieve certificate"
// @Router /certificates/{hash} [get]
func (h *Handler) GetByHash(c *gin.Context) {
	hash := c.Param("hash")

	// Validate hash format (0x + 64 hex characters = 66 total)
	if len(hash) != 66 || hash[:2] != "0x" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid hash format",
			"message": "Hash must be 66 characters (0x + 64 hex digits)",
		})
		return
	}

	cert, err := blockchain.VerifyCertificate(h.Client, h.Config.ContractAddress, hash)
	if err != nil {
		log.Printf("Error retrieving certificate: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve certificate"})
		return
	}

	if !cert.IsValid {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "not_found",
			"message": "Certificate not found or invalid",
		})
		return
	}

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
