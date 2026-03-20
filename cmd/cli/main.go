package main

// export GOTMPDIR=~/go-cache/tmp
import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
	"github.com/joho/godotenv"
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
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to Alchemy
	client, err := ethclient.Dial(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		log.Fatalf("Cant Connect to Alchemy: %v", err)
	}
	fmt.Println("Success connecting to Alchemy")

	// Generate hash of the PDF file
	hash, err := generateHash(pdfPath)
	if err != nil {
		log.Fatalf("Error generating hash: %v", err)
	}

	fmt.Printf("Generated Hash: %s\n", hash)

	// Action based on command line argument
	switch action {
	case "register":
		registerCertificate(client, common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), hash)
	case "verify":
		verifyCertificate(client, common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), hash)
	default:
		fmt.Println("Unknown action. Use 'register' to register the certificate or 'verify' to verify it.")
	}
}

func generateHash(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := crypto.Keccak256Hash(file).Hex()

	return hash, nil
}

func registerCertificate(client *ethclient.Client, contractAddress common.Address, fileHash string) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Invalid PRIVATE_KEY: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Instantiate the smart contract
	instance, err := blockchain.NewCertifyer(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// Register the Hash
	// Convert the hash to bytes32
	var dataHash [32]byte
	copy(dataHash[:], common.FromHex(fileHash))

	tx, err := instance.RegisterCertificate(auth, dataHash)
	if err != nil {
		log.Fatalf("Failed to register hash: %v", err)
	}

	fmt.Printf("Hash registered successfully! Transaction Hash: %s\n", tx.Hash().Hex())

}

func verifyCertificate(client *ethclient.Client, contractAddress common.Address, fileHash string) {
	instance, err := blockchain.NewCertifyer(contractAddress, client)

	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	var hash [32]byte
	hashBytes, _ := hex.DecodeString(fileHash[2:]) // Remove "0x" prefix
	copy(hash[:], hashBytes)

	// Free call to check if the hash is registered
	isValid, err := instance.Certificates(nil, hash)
	if err != nil {
		log.Fatalf("Failed to verify hash: %v", err)
	}

	fmt.Println("\n---Verification Result---")
	if isValid {
		fmt.Printf("The certificate with hash %s is valid and registered on the blockchain.\n", fileHash)
	} else {
		fmt.Printf("The certificate with hash %s is NOT registered on the blockchain.\n", fileHash)
	}
	fmt.Println("-------------------------")
}
