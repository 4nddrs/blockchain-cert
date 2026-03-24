package main

// export GOTMPDIR=~/go-cache/tmp
import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
)

func main() {

	// Check command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}
	action := os.Args[1]
	pdfPath := os.Args[2]

	// Load environment variables from .env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	url := os.Getenv("ALCHEMY_URL")
	if url == "" {
		log.Fatal("Error: Cant found ALCHEMY_URL in .env file")
	}

	// Connect to Alchemy
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		log.Fatalf("Cant Connect to Alchemy: %v", err)
	}
	fmt.Println("Success connecting to Alchemy")

	// Generate hash of the PDF file
	hash, err := blockchain.GenerateHash(pdfPath)
	if err != nil {
		log.Fatalf("Error generating hash: %v", err)
	}

	fmt.Printf("Generated Hash: %s\n", hash)

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	// Action based on command line argument
	switch action {
	case "register":

		// Validate arguments
		if len(os.Args) < 6 {
			fmt.Println("Usage: go run main.go register <file.pdf> <studentName> <CourseName> <IssuerName>")
		}

		studentName := os.Args[3]
		courseName := os.Args[4]
		issuerName := os.Args[5]
		privateKey := os.Getenv("PRIVATE_KEY")

		txHash, err := blockchain.RegisterCertificate(client, privateKey, contractAddress, hash, studentName, courseName, issuerName)
		if err != nil {
			log.Fatalf("Error registering certificate: %v", err)
		}

		fmt.Printf("Certificate registered successfully! Transaction Hash: %s\n", txHash)
		fmt.Printf("Name: %s\nCourse: %s\nIssuer: %s\n", studentName, courseName, issuerName)

	case "verify":
		cert, err := blockchain.VerifyCertificate(client, contractAddress, hash)

		if err != nil {
			log.Fatalf("Error verifying certificate: %v", err)
		}

		fmt.Printf("Certificate Details:\n")
		if cert.IsValid {
			fmt.Printf("Status: Valid\n")
			fmt.Printf("Student Name: %s\n", cert.StudentName)
			fmt.Printf("Course Name: %s\n", cert.CourseName)
			fmt.Printf("Issuer Name: %s\n", cert.IssuerName)
			fmt.Printf("Registration Date: %s\n", cert.DateEmited.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("Status: Invalid\n")
		}

	default:
		fmt.Println("Unknown action. Use 'register' to register the certificate or 'verify' to verify it.")
	}
}
