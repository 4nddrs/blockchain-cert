package main

// export GOTMPDIR=~/go-cache/tmp
import (
	"fmt"
	"log"
	"os"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
	"github.com/4nddrs/blockchain-cert/internal/crypto"
	"github.com/joho/godotenv"
)

func main() {

	// 1. Load environment variables from .env file
	err := godotenv.Load("../../.env")
	url := os.Getenv("ALCHEMY_URL")
	if url == "" {
		log.Fatal("Error: ALCHEMY_URL no encontrada en el .env")
	}
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 2. Generate file Hash
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}
	pdfPath := os.Args[1]
	hash, _ := crypto.GenerateFileHash(pdfPath)
	fmt.Printf("Generated Hash: %s\n", hash)

	// 3. Connect to Alchemy
	client, err := blockchain.Connect(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		log.Fatalf("Cant Connect to Alchemy", err)
	}
	fmt.Println("Success connecting to Alchemy")
	_ = client
}
