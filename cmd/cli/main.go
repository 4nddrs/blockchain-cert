package main

// export GOTMPDIR=~/go-cache/tmp
import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/4nddrs/blockchain-cert/internal/blockchain"
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
	hashBytes, err := os.ReadFile(pdfPath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	hash := crypto.Keccak256Hash(hashBytes).Hex()
	fmt.Printf("Generated Hash: %s\n", hash)

	// 3. Connect to Alchemy
	client, err := blockchain.Connect(os.Getenv("ALCHEMY_URL"))
	if err != nil {
		log.Fatalf("Cant Connect to Alchemy: %v", err)
	}
	fmt.Println("Success connecting to Alchemy")
	_ = client

	// 4. Configure the signature (Auth)
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

	// 5. Instantiate the smart contract
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := blockchain.NewCertifyer(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// 6. Register the Hash
	// Convert the hash to bytes32
	var dataHash [32]byte
	copy(dataHash[:], common.FromHex(hash))

	tx, err := instance.RegisterCertificate(auth, dataHash)
	if err != nil {
		log.Fatalf("Failed to register hash: %v", err)
	}

	fmt.Printf("Hash registered successfully! Transaction Hash: %s\n", tx.Hash().Hex())
}
