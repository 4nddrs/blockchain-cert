package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	AlchemyURL      string
	ContractAddress common.Address
	PrivateKey      string
	AdminSecret     string
	ServerPort      string
	CORSOrigins     []string
	TempUploadDir   string
}

// findEnvFile searches for .env file starting from current directory up to root
func findEnvFile() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return ""
}

// Load reads environment variables and returns a Config instance
func Load() (*Config, error) {
	// Try to find and load .env file (optional in production)
	if envPath := findEnvFile(); envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("Warning: found .env at %s but failed to load: %v\n", envPath, err)
		} else {
			log.Printf("Loaded .env from: %s\n", envPath)
		}
	} else {
		log.Println("Warning: .env file not found in current or parent directories, using environment variables")
	}

	alchemyURL := os.Getenv("ALCHEMY_URL")
	if alchemyURL == "" {
		return nil, fmt.Errorf("ALCHEMY_URL environment variable is required")
	}

	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddr == "" {
		return nil, fmt.Errorf("CONTRACT_ADDRESS environment variable is required")
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		return nil, fmt.Errorf("PRIVATE_KEY environment variable is required")
	}

	adminSecret := os.Getenv("ADMIN_SECRET")
	if adminSecret == "" {
		log.Println("Warning: ADMIN_SECRET not set, admin endpoints will be disabled")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins == "" {
		corsOrigins = "http://localhost:3000"
	}

	tempDir := os.Getenv("TEMP_UPLOAD_DIR")
	if tempDir == "" {
		tempDir = "./temp_uploads"
	}

	// Create temp directory if it doesn't exist
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	return &Config{
		AlchemyURL:      alchemyURL,
		ContractAddress: common.HexToAddress(contractAddr),
		PrivateKey:      privateKey,
		AdminSecret:     adminSecret,
		ServerPort:      serverPort,
		CORSOrigins:     []string{corsOrigins},
		TempUploadDir:   tempDir,
	}, nil
}
