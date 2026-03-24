package blockchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CertificateData struct {
	IsValid     bool
	StudentName string
	CourseName  string
	IssuerName  string
	DateEmited  time.Time
}

// Function to generate a hash of the certificate file
func GenerateHash(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := crypto.Keccak256Hash(file).Hex()

	return hash, nil
}

func RegisterCertificate(
	client *ethclient.Client,
	privateKeyHex string,
	contractAddress common.Address,
	fileHash string,
	studentName string,
	courseName string,
	issuerName string) (string, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("Failed to parse private key: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("Failed to create transactor: %v", err)
	}

	// Instantiate the smart contract
	instance, err := NewCertifyer(contractAddress, client)
	if err != nil {
		return "", fmt.Errorf("Failed to instantiate contract: %v", err)
	}

	// Register the Hash
	// Convert the hash to bytes32
	var dataHash [32]byte
	copy(dataHash[:], common.FromHex(fileHash))

	tx, err := instance.RegisterCertificate(auth, dataHash, studentName, courseName, issuerName)
	if err != nil {
		return "", fmt.Errorf("Failed to register certificate: %v", err)
	}

	return tx.Hash().Hex(), nil
}

// Verify if the certificate hash is registered on the blockchain
func VerifyCertificate(
	client *ethclient.Client,
	contractAddress common.Address,
	fileHash string) (*CertificateData, error) {

	instance, err := NewCertifyer(contractAddress, client)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate contract: %v", err)
	}

	var hash [32]byte
	hashBytes, _ := hex.DecodeString(fileHash[2:]) // Remove "0x" prefix
	copy(hash[:], hashBytes)

	// Free call to check if the hash is registered
	cert, err := instance.Certificates(nil, hash)
	if err != nil {
		return nil, fmt.Errorf("Failed to verify certificate: %v", err)
	}

	return &CertificateData{
		IsValid:     cert.IsValid,
		StudentName: cert.StudentName,
		CourseName:  cert.CourseName,
		IssuerName:  cert.IssuerName,
		DateEmited:  time.Unix(cert.DateEmited.Int64(), 0),
	}, nil
}
