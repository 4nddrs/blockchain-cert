# Blockchain Certificate Verification System

A professional-grade blockchain-based document certification system built with Go and Ethereum smart contracts. This system provides immutable proof of document authenticity by registering cryptographic hashes on the Polygon blockchain.

## 🎯 Purpose & Importance

### The Problem

Traditional document certification systems rely on centralized authorities that can be compromised, corrupted, or become unavailable. Physical seals and signatures can be forged, and centralized databases can be tampered with.

### The Solution

This system creates an **immutable, decentralized registry** of document authenticity. Instead of storing the actual files (which would be expensive and impractical on blockchain), we store their cryptographic "fingerprints" (hashes). This approach provides:

- **Immutability**: Once registered, a certificate cannot be altered or deleted
- **Decentralization**: No single point of failure or control
- **Privacy**: The actual documents remain with their owners
- **Public Verifiability**: Anyone can verify a document's authenticity
- **Cryptographic Security**: SHA-256 hashing ensures document integrity
- **Cost Efficiency**: Only hashes are stored on-chain, minimizing gas costs

### Real-World Applications

- **Academic Credentials**: Universities can certify degrees and transcripts
- **Legal Documents**: Lawyers can timestamp and verify contracts
- **Medical Records**: Hospitals can prove authenticity of patient records
- **Intellectual Property**: Creators can establish proof of creation dates
- **Supply Chain**: Companies can verify product authenticity certificates

---

## 🏗️ Architecture

### System Components

```
┌─────────────────────────────────────────────────────────────────┐
│                         USER INTERACTION                        │
│                    (CLI - Command Line Interface)               │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      APPLICATION LAYER (Go)                     |
│  ┌──────────────┐  ┌──────────────┐  ┌────────────────────┐     │
│  │ Crypto Layer │  │   Ethereum   │  │   Smart Contract   │     │
│  │  (Hashing)   │  │    Client    │  │    Bindings (ABI)  │     │
│  └──────────────┘  └──────────────┘  └────────────────────┘     │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                   INFRASTRUCTURE LAYER                          │
│                  (Alchemy RPC Node Provider)                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      BLOCKCHAIN LAYER                           │
│              Polygon Amoy Testnet (Ethereum-compatible)         │
│                   Smart Contract: Certifyer.sol                 │
└─────────────────────────────────────────────────────────────────┘
```

### Data Flow

```
Step 1: Local Processing (Offline & Private)
┌────────────┐
│ PDF/File   │──▶ SHA-256 Hash ──▶ 0xabcd1234... (32 bytes)
└────────────┘

Step 2: Digital Signature (ECDSA)
┌──────────────┐
│ Private Key  │──▶ Sign Transaction ──▶ Authorized Transaction
└──────────────┘

Step 3: Blockchain Submission
┌────────────────┐
│ Go CLI         │──▶ Alchemy (RPC) ──▶ Polygon Network
└────────────────┘

Step 4: Immutable Storage
┌─────────────────┐
│ Smart Contract  │──▶ mapping(hash => true) ──▶ Event Emitted
└─────────────────┘

Step 5: Public Verification
Anyone can query: validateCertificate(hash) ──▶ true/false
```

---

## 🔧 Technology Stack

| Layer                      | Technology           | Purpose                                 |
| -------------------------- | -------------------- | --------------------------------------- |
| **Language**               | Go 1.25+             | High-performance CLI application        |
| **Smart Contract**         | Solidity 0.8.19      | On-chain certificate registry           |
| **Blockchain**             | Polygon Amoy Testnet | Ethereum-compatible network             |
| **Node Provider**          | Alchemy              | RPC gateway to blockchain               |
| **Development Framework**  | Foundry (Forge/Cast) | Smart contract compilation & deployment |
| **Blockchain Client**      | go-ethereum v1.17.1  | Ethereum interaction library            |
| **Cryptography**           | SHA-256/Keccak256    | Document hashing                        |
| **Environment Management** | godotenv             | Secure configuration                    |
| **JSON Processing**        | jq                   | ABI extraction                          |
| **Code Generation**        | abigen               | Go bindings from Solidity ABI           |

---

## 📁 Project Structure

```
blockchain-cert/
├── cmd/
│   └── cli/
│       ├── main.go           # CLI entry point with clean code architecture
│       └── title.pdf          # Example test file
├── internal/
│   ├── blockchain/
│   │   └── certifyer.go       # Auto-generated contract bindings
├── contracts/
│   └── Certifyer.sol          # Solidity smart contract
├── out/                       # Foundry build artifacts (auto-generated)
│   └── Certifyer.sol/
│       └── Certifyer.json     # Compiled contract ABI + bytecode
├── .env                       # Environment variables (NOT committed)
├── .gitignore
├── go.mod                     # Go module definition
├── go.sum                     # Go dependencies lockfile
├── AGENTS.md                  # Developer documentation
└── README.md                  # This file
```

### Key Directories

- **`cmd/cli/`**: Command-line interface executable with register and verify functionality
- **`internal/blockchain/`**: Smart contract bindings for Ethereum interaction
- **`contracts/`**: Solidity smart contracts source code
- **`out/`**: Foundry compilation artifacts

---

## 🚀 Installation & Setup

### Prerequisites

1. **Go 1.25+**: [Install Go](https://golang.org/doc/install)
2. **Foundry**: Solidity development toolkit
3. **jq**: JSON processor
4. **Alchemy Account**: Free RPC node access
5. **Metamask Wallet**: For private key and testnet tokens

### Step 1: Install System Dependencies

```bash
# Install Foundry (Rust-based Solidity toolkit)
curl -L https://foundry.paradigm.xyz | bash
foundryup

# Install jq (JSON processor)
sudo apt install jq  # Debian/Ubuntu
# or
brew install jq      # macOS
```

### Step 2: Clone & Install Dependencies

```bash
git clone <your-repo-url>
cd blockchain-cert

# Install Go dependencies
go mod download
go mod tidy

# Install abigen (Go bindings generator)
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:$(go env GOPATH)/bin
```

### Step 3: Set Up Alchemy

1. Create account at [alchemy.com](https://www.alchemy.com/)
2. Create a new app:
   - **Chain**: Polygon
   - **Network**: Amoy (Testnet)
3. Copy the **HTTPS RPC URL**

### Step 4: Configure Metamask

1. Add Polygon Amoy testnet:
   - Visit [Chainlist](https://chainlist.org/)
   - Search "Polygon Amoy"
   - Connect Metamask
2. Get testnet tokens:
   - Visit [Polygon Faucet](https://faucet.polygon.technology/)
   - Select Amoy network
   - Enter your wallet address
   - Request POL tokens

### Step 5: Create Environment File

Create `.env` in the project root:

```env
# Alchemy RPC endpoint
ALCHEMY_URL=https://polygon-amoy.g.alchemy.com/v2/YOUR_API_KEY

# Your Metamask private key (starts with 0x)
# ⚠️ NEVER commit this file to git
PRIVATE_KEY=0xYOUR_PRIVATE_KEY_HERE

# Contract address (fill after deployment)
CONTRACT_ADDRESS=0x
```

**⚠️ Security Warning**: Never commit `.env` to version control. Ensure it's listed in `.gitignore`.

---

## 📜 Smart Contract Deployment

### Step 1: Compile the Contract

```bash
# Compile Solidity code
forge build

# Output: ./out/Certifyer.sol/Certifyer.json
```

### Step 2: Verify Wallet Balance

```bash
# Check you have testnet POL tokens
cast balance YOUR_WALLET_ADDRESS --rpc-url $ALCHEMY_URL
```

### Step 3: Deploy to Polygon Amoy

```bash
# Deploy the contract
forge create --rpc-url $ALCHEMY_URL \
             --private-key $PRIVATE_KEY \
             contracts/Certifyer.sol:Certifyer \
             --broadcast
```

**Expected Output:**

```
Deployer: 0xYourAddress
Deployed to: 0xCONTRACT_ADDRESS_HERE
Transaction hash: 0xTX_HASH
```

### Step 4: Update Configuration

Copy the deployed contract address to `.env`:

```env
CONTRACT_ADDRESS=0xYOUR_DEPLOYED_CONTRACT_ADDRESS
```

### Step 5: Generate Go Bindings

```bash
# Extract ABI from Foundry output
jq .abi out/Certifyer.sol/Certifyer.json > out/Certifyer.abi

# Generate Go bindings
abigen --abi out/Certifyer.abi \
       --pkg blockchain \
       --type Certifyer \
       --out internal/blockchain/certifyer.go
```

---

## 🎮 Usage

The CLI now supports two main operations: **registering** and **verifying** certificates.

### Register a Certificate

```bash
cd cmd/cli

# Register a document certificate on the blockchain
go run main.go register <file_path>

# Example
go run main.go register title.pdf
```

**Expected Output:**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
Hash registered successfully! Transaction Hash: 0x123456789...
```

### Verify a Certificate

```bash
cd cmd/cli

# Verify if a document is registered on the blockchain
go run main.go verify <file_path>

# Example
go run main.go verify title.pdf
```

**Expected Output (if registered):**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab

---Verification Result---
The certificate with hash 0xabcd1234... is valid and registered on the blockchain.
-------------------------
```

**Expected Output (if NOT registered):**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab

---Verification Result---
The certificate with hash 0xabcd1234... is NOT registered on the blockchain.
-------------------------
```

### Verify Transaction on Blockchain Explorer

1. Copy the transaction hash from registration
2. Visit [Polygon Amoy Explorer](https://amoy.polygonscan.com/)
3. Paste the transaction hash
4. View the `CertificateCreated` event

### Query Contract State via CLI

```bash
# Alternative: Check if a hash is certified using Foundry
cast call $CONTRACT_ADDRESS \
  "certificates(bytes32)(bool)" \
  0xYOUR_HASH_HERE \
  --rpc-url $ALCHEMY_URL
```

---

## 🔒 Security Considerations

### Current Implementation

✅ **Strengths:**

- Private keys stored in `.env` (not hardcoded)
- Admin-only certificate registration (access control)
- Immutable records (cannot be deleted or modified)
- Cryptographic hashing (SHA-256/Keccak256)

⚠️ **Areas for Improvement:**

- No private key encryption at rest
- Single admin point of failure
- No rate limiting on RPC calls
- No input validation on file paths
- API keys visible in plain text `.env`

### Best Practices

1. **Never commit `.env` to version control**
2. **Use hardware wallets for production**
3. **Implement multi-signature admin control**
4. **Validate all file inputs before hashing**
5. **Use secret management services (AWS Secrets Manager, HashiCorp Vault)**
6. **Monitor Alchemy API rate limits**
7. **Implement audit logging for all certifications**

---

## 🧪 Development Workflow

### Typical Development Flow

```bash
# 1. Make changes to Go code
vim cmd/cli/main.go

# 2. Test registration locally
cd cmd/cli
go run main.go register test_file.pdf

# 3. Verify the certificate
go run main.go verify test_file.pdf

# 4. Check transaction on blockchain explorer
# Visit https://amoy.polygonscan.com/ with transaction hash
```

### Modifying the Smart Contract

```bash
# 1. Edit contract
vim contracts/Certifyer.sol

# 2. Compile
forge build

# 3. Deploy new version
forge create --rpc-url $ALCHEMY_URL \
             --private-key $PRIVATE_KEY \
             contracts/Certifyer.sol:Certifyer \
             --broadcast

# 4. Update CONTRACT_ADDRESS in .env

# 5. Regenerate Go bindings
jq .abi out/Certifyer.sol/Certifyer.json > out/Certifyer.abi
abigen --abi out/Certifyer.abi \
       --pkg blockchain \
       --type Certifyer \
       --out internal/blockchain/certifyer.go

# 6. Rebuild Go app
cd cmd/cli
go build
```

---

## 🏛️ Clean Code Architecture

The CLI has been refactored following clean code principles:

### Separation of Concerns

The main file is now organized into **focused functions** with single responsibilities:

- **`main()`**: Orchestrates flow, parses arguments, delegates to specialized functions
- **`generateHash()`**: Pure function for file hashing
- **`registerCertificate()`**: Handles blockchain registration logic
- **`verifyCertificate()`**: Handles blockchain verification logic

### Benefits of Current Architecture

✅ **Modularity**: Each function does one thing well  
✅ **Testability**: Functions can be unit tested independently  
✅ **Readability**: Clear function names describe intent  
✅ **Maintainability**: Easy to modify or extend functionality  
✅ **Error Handling**: Consistent error propagation pattern  

### Command Pattern

The CLI uses a simple command pattern:

```go
switch action {
case "register":
    registerCertificate(client, contractAddress, hash)
case "verify":
    verifyCertificate(client, contractAddress, hash)
default:
    fmt.Println("Unknown action. Use 'register' or 'verify'.")
}
```

This makes it easy to add new commands in the future (e.g., `batch`, `revoke`, `list`).

---

## 🛠️ Technical Deep Dive

### How Hashing Works

The system uses a clean, modular approach to file hashing:

```go
// generateHash creates a Keccak256 hash of any file
func generateHash(filePath string) (string, error) {
    file, err := os.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    hash := crypto.Keccak256Hash(file).Hex()
    return hash, nil
}
// Result: 0x + 64 hex characters (32 bytes)
```

**Why Keccak256?**

- Ethereum standard (compatible with Solidity `bytes32`)
- Deterministic (same file = same hash)
- Collision-resistant (practically impossible to forge)
- One-way function (cannot reverse hash to file)

### Smart Contract Logic

```solidity
// Admin-controlled registry
address public admin;
mapping(bytes32 => bool) public certificates;

// Only admin can certify
function registerCertificate(bytes32 datahash) public {
    require(msg.sender == admin, "Only the admin can certify documents");
    certificates[datahash] = true;
    emit CertificateCreated(datahash, block.timestamp);
}

// Public mapping for verification (anyone can check)
mapping(bytes32 => bool) public certificates;
```

**Key Design Decisions:**

1. **Mapping over array**: O(1) lookup time
2. **Events for indexing**: Off-chain services can listen for new certificates
3. **Public mapping**: Verification costs no gas (automatic getter)
4. **No deletion**: Immutable by design

### Certificate Registration Flow

```go
func registerCertificate(client *ethclient.Client, contractAddress common.Address, fileHash string) {
    // 1. Load private key and create authorized signer
    privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    chainID, _ := client.NetworkID(context.Background())
    auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    
    // 2. Instantiate smart contract
    instance, _ := blockchain.NewCertifyer(contractAddress, client)
    
    // 3. Convert hash to bytes32
    var dataHash [32]byte
    copy(dataHash[:], common.FromHex(fileHash))
    
    // 4. Submit transaction
    tx, _ := instance.RegisterCertificate(auth, dataHash)
    fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
}
```

### Certificate Verification Flow

```go
func verifyCertificate(client *ethclient.Client, contractAddress common.Address, fileHash string) {
    // 1. Instantiate smart contract (no auth needed for read-only)
    instance, _ := blockchain.NewCertifyer(contractAddress, client)
    
    // 2. Convert hash to bytes32
    var hash [32]byte
    hashBytes, _ := hex.DecodeString(fileHash[2:]) // Remove "0x" prefix
    copy(hash[:], hashBytes)
    
    // 3. Free call to check registration (no gas cost)
    isValid, _ := instance.Certificates(nil, hash)
    
    // 4. Display result
    if isValid {
        fmt.Println("Certificate is valid and registered")
    } else {
        fmt.Println("Certificate is NOT registered")
    }
}
```

### Transaction Signing (ECDSA)

```go
// 1. Load private key
privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

// 2. Get network ID
chainID, _ := client.NetworkID(context.Background())

// 3. Create authorized signer
auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

// 4. Send transaction
tx, _ := instance.RegisterCertificate(auth, dataHash)
```

This proves **you** authorized the transaction (non-repudiation).

---

## 🔍 Troubleshooting

### Common Issues

#### "Usage: go run main.go <file_path>"

**Old error**: You forgot to specify the action.

**New usage**:
```bash
go run main.go register <file_path>  # To register
go run main.go verify <file_path>    # To verify
```

#### "Error: Cant found ALCHEMY_URL in .env file"

- Verify `.env` exists in project root
- Check `ALCHEMY_URL` is set and not commented
- Ensure running from correct directory (`cmd/cli/`)

#### "Cant Connect to Alchemy"

- Verify Alchemy URL is correct
- Check internet connectivity
- Confirm API key is valid
- Test with: `curl $ALCHEMY_URL -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'`

#### "Failed to register hash: insufficient funds"

- Request testnet tokens from [Polygon Faucet](https://faucet.polygon.technology/)
- Check balance: `cast balance YOUR_ADDRESS --rpc-url $ALCHEMY_URL`

#### "Invalid PRIVATE_KEY"

- Ensure key starts with `0x`
- Verify key is 64 hex characters (32 bytes)
- Export from Metamask: Account Details → Export Private Key

#### "Unknown action. Use 'register' or 'verify'."

- You provided an invalid action
- Valid actions: `register`, `verify`
- Example: `go run main.go register file.pdf`

#### Go module errors

```bash
go clean -modcache
go mod download
go mod tidy
```

---

## 📊 Performance Metrics

| Operation              | Time  | Cost (Gas)            |
| ---------------------- | ----- | --------------------- |
| Local hash generation  | ~10ms | Free                  |
| Transaction submission | 2-30s | ~50,000 gas (~$0.01)  |
| Certificate validation | <1s   | Free (view function)  |
| Contract deployment    | ~30s  | ~300,000 gas (~$0.05) |

_Gas costs on Polygon Amoy testnet (free). Mainnet costs will vary with POL price._

---

## 🚧 Future Development

### Planned Features

1. ✅ Certificate registration
2. ✅ Certificate validation CLI command
3. ⏳ Batch certificate registration
4. ⏳ Certificate revocation mechanism
5. ⏳ Web interface (REST API)
6. ⏳ Multi-signature admin control
7. ⏳ IPFS integration for metadata storage
8. ⏳ Email notifications on certification
9. ⏳ QR code generation for verification
10. ⏳ Automated testing suite

### Roadmap to Production

- [ ] Add comprehensive unit tests
- [ ] Implement integration tests with mock blockchain
- [ ] Set up CI/CD pipeline
- [ ] Add input validation and sanitization
- [ ] Implement rate limiting
- [ ] Add structured logging (zerolog/zap)
- [ ] Create deployment scripts for mainnet
- [ ] Set up monitoring and alerting
- [ ] Write API documentation (OpenAPI/Swagger)
- [ ] Implement backup admin mechanism (multi-sig)

---

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Follow existing code patterns and conventions
4. Add tests for new functionality
5. Ensure all tests pass (`go test ./...`)
6. Commit with clear messages
7. Push to your fork
8. Open a Pull Request

### Code Style

- Use `gofmt` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Add comments for exported functions
- Keep functions small and focused

---

## 📄 License

MIT License - See LICENSE file for details

---

## 🔗 Useful Resources

- [Polygon Documentation](https://docs.polygon.technology/)
- [Foundry Book](https://book.getfoundry.sh/)
- [go-ethereum Documentation](https://geth.ethereum.org/docs)
- [Solidity Docs](https://docs.soliditylang.org/)
- [Alchemy Documentation](https://docs.alchemy.com/)

---

## 📞 Support

For questions or issues:

- Open an issue on GitHub
- Check existing documentation in `AGENTS.md`
- Review transaction on [Polygon Amoy Explorer](https://amoy.polygonscan.com/)

---

## 🙏 Acknowledgments

Built with:

- **Go** - The Go Authors
- **Solidity** - Ethereum Foundation
- **Foundry** - Paradigm
- **Polygon** - Polygon Labs
- **Alchemy** - Alchemy Insights Inc.

---

_Built by engineers, for engineers. Immutable truth on the blockchain._
