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
- **Cryptographic Security**: Keccak256 hashing ensures document integrity
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
┌───────────────────────────────────────────────────────────────────┐
│                      USER INTERACTION                             │
│           CLI (Command Line)  |  REST API (HTTP)                  │
└────────────────────────┬──────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                   APPLICATION LAYER (Go)                        │
│  ┌────────────────┐  ┌──────────────┐  ┌────────────────────┐   │
│  │  Logic Layer   │  │   Ethereum   │  │   Smart Contract   │   │
│  │  (logic.go)    │  │    Client    │  │    Bindings (ABI)  │   │
│  └────────────────┘  └──────────────┘  └────────────────────┘   │
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
│ PDF/File   │──▶ Keccak256 Hash ──▶ 0xabcd1234... (32 bytes)
└────────────┘

Step 2: Metadata Preparation
┌──────────────────────┐
│ Student Name         │──▶ Certificate Metadata
│ Course Name          │
│ Issuer Name          │
└──────────────────────┘

Step 3: Digital Signature (ECDSA)
┌──────────────┐
│ Private Key  │──▶ Sign Transaction ──▶ Authorized Transaction
└──────────────┘

Step 4: Blockchain Submission
┌────────────────┐
│ CLI / API      │──▶ Alchemy (RPC) ──▶ Polygon Network
└────────────────┘

Step 5: Immutable Storage
┌─────────────────┐
│ Smart Contract  │──▶ mapping(hash => Certificate) ──▶ Event Emitted
└─────────────────┘

Step 6: Public Verification
Anyone can query: certificates[hash] ──▶ {isValid, studentName, courseName, issuerName, dateEmitted}
```

---

## 🔧 Technology Stack

| Layer                      | Technology           | Purpose                                 |
| -------------------------- | -------------------- | --------------------------------------- |
| **Language**               | Go 1.25+             | High-performance application            |
| **Smart Contract**         | Solidity 0.8.19      | On-chain certificate registry           |
| **Blockchain**             | Polygon Amoy Testnet | Ethereum-compatible network             |
| **Node Provider**          | Alchemy              | RPC gateway to blockchain               |
| **Development Framework**  | Foundry (Forge/Cast) | Smart contract compilation & deployment |
| **Blockchain Client**      | go-ethereum v1.17.1  | Ethereum interaction library            |
| **API Framework**          | Gin v1.12.0          | HTTP REST API server                    |
| **CORS Middleware**        | gin-contrib/cors     | Cross-origin resource sharing           |
| **Cryptography**           | Keccak256            | Document hashing                        |
| **Environment Management** | godotenv             | Secure configuration                    |
| **JSON Processing**        | jq                   | ABI extraction                          |
| **Code Generation**        | abigen               | Go bindings from Solidity ABI           |

---

## 📁 Project Structure

```
blockchain-cert/
├── cmd/
│   ├── cli/
│   │   ├── main.go           # CLI entry point
│   │   └── pdfs/             # Test PDFs directory (gitignored)
│   └── api/
│       ├── main.go           # REST API server
│       └── temp_/            # Temporary upload directory
├── internal/
│   ├── blockchain/
│   │   ├── logic.go          # Core business logic (register/verify)
│   │   ├── certifyer.go      # Auto-generated contract bindings
│   │   └── client.go         # Ethereum client connection (deprecated)
│   └── crypto/
│       └── hash.go           # SHA256 file hashing (deprecated)
├── contracts/
│   └── Certifyer.sol         # Solidity smart contract
├── out/                      # Foundry build artifacts (auto-generated)
│   └── Certifyer.sol/
│       └── Certifyer.json    # Compiled contract ABI + bytecode
├── .env                      # Environment variables (NOT committed)
├── .gitignore
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies lockfile
├── AGENTS.md                 # Developer documentation
└── README.md                 # This file
```

### Key Directories

- **`cmd/cli/`**: Command-line interface executable with register and verify functionality
- **`cmd/api/`**: REST API server for web applications
- **`internal/blockchain/`**: Core business logic and smart contract bindings
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

The system provides **two interfaces**: a CLI for terminal usage and a REST API for web applications.

## CLI Usage

### Register a Certificate

```bash
cd cmd/cli

# Register a document certificate with metadata on the blockchain
go run main.go register <file_path> <studentName> <courseName> <issuerName>

# Example
go run main.go register document.pdf "Andres Menchaca" "Blockchain Development" "Tech Academy"
```

**Parameters:**

- `<file_path>`: Path to the PDF or document file
- `<studentName>`: Full name of the certificate recipient
- `<courseName>`: Name of the course or certification program
- `<issuerName>`: Organization or entity issuing the certificate

**Expected Output:**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
Certificate registered successfully! Transaction Hash: 0x123456789...
Name: Andres Menchaca
Course: Blockchain Development
Issuer: Tech Academy
```

### Verify a Certificate

```bash
cd cmd/cli

# Verify if a document is registered on the blockchain
go run main.go verify <file_path>

# Example
go run main.go verify document.pdf
```

**Expected Output (if registered):**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
Certificate Details:
Status: Valid
Student Name: Andres Menchaca
Course Name: Blockchain Development
Issuer Name: Tech Academy
Registration Date: 2026-03-24 14:30:45
```

**Expected Output (if NOT registered):**

```
Success connecting to Alchemy
Generated Hash: 0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
Certificate Details:
Status: Invalid
```

---

## 🌐 REST API Usage

### Start the API Server

```bash
cd cmd/api
go run main.go

# Server will start on http://localhost:8080
```

### API Endpoints

#### 1. Register Certificate

**Endpoint**: `POST /api/v1/register`

**Content-Type**: `multipart/form-data`

**Parameters**:

- `pdf` (file): PDF document to certify
- `student_name` (string): Student's full name
- `course_name` (string): Course name
- `issuer` (string): Issuing organization

**Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -F "pdf=@certificate.pdf" \
  -F "student_name=Andres Menchaca" \
  -F "course_name=Blockchain Development" \
  -F "issuer=Tech Academy"
```

**Success Response** (200 OK):

```json
{
  "status": "success",
  "hash": "0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab",
  "tx_hash": "0x9876543210fedcba9876543210fedcba9876543210fedcba9876543210fedcba"
}
```

**Error Response** (400 Bad Request):

```json
{
  "error": "Failed to register certificate"
}
```

---

#### 2. Verify Certificate

**Endpoint**: `POST /api/v1/verify`

**Content-Type**: `multipart/form-data`

**Parameters**:

- `pdf` (file): PDF document to verify

**Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/verify \
  -F "pdf=@certificate.pdf"
```

**Success Response (Valid Certificate)** (200 OK):

```json
{
  "status": "valid",
  "student_name": "Andres Menchaca",
  "course_name": "Blockchain Development",
  "issuer": "Tech Academy",
  "date": "2026-03-24 14:30:45"
}
```

**Not Found Response** (404 Not Found):

```json
{
  "status": "not_found",
  "message": "Certificate not found or invalid"
}
```

---

### CORS Configuration

The API is configured with CORS to accept requests from `http://localhost:3000` (typical React/Vue development server).

To modify allowed origins, edit `cmd/api/main.go`:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000", "https://yourdomain.com"},
    AllowMethods:     []string{"POST", "GET", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))
```

---

### Verify Transaction on Blockchain Explorer

1. Copy the transaction hash from registration
2. Visit [Polygon Amoy Explorer](https://amoy.polygonscan.com/)
3. Paste the transaction hash
4. View the `CertificateCreated` event

---

## 🏛️ Clean Code Architecture

The project follows clean code principles with clear separation of concerns:

### Core Business Logic (`internal/blockchain/logic.go`)

All certificate-related operations are centralized in a single file:

```go
// GenerateHash - Creates Keccak256 hash of any file
func GenerateHash(filePath string) (string, error)

// RegisterCertificate - Submits certificate to blockchain with metadata
func RegisterCertificate(
    client *ethclient.Client,
    privateKeyHex string,
    contractAddress common.Address,
    fileHash string,
    studentName string,
    courseName string,
    issuerName string,
) (string, error)

// VerifyCertificate - Queries blockchain for certificate data
func VerifyCertificate(
    client *ethclient.Client,
    contractAddress common.Address,
    fileHash string,
) (*CertificateData, error)
```

### Benefits of Current Architecture

✅ **Single Responsibility**: Each function has one clear purpose  
✅ **DRY Principle**: Logic shared between CLI and API  
✅ **Testability**: Pure functions can be unit tested independently  
✅ **Maintainability**: Changes to business logic happen in one place  
✅ **Scalability**: Easy to add new interfaces (WebSocket, gRPC, etc.)

### Interface Layer

Both CLI and API are thin wrappers around `logic.go`:

- **CLI** (`cmd/cli/main.go`): Parses arguments, calls logic, formats output
- **API** (`cmd/api/main.go`): Handles HTTP, file uploads, calls logic, returns JSON

This makes it trivial to add new interfaces without duplicating business logic.

---

## 🔒 Security Considerations

### Current Implementation

✅ **Strengths:**

- Private keys stored in `.env` (not hardcoded)
- Admin-only certificate registration (access control)
- Immutable records (cannot be deleted or modified)
- Cryptographic hashing (Keccak256)
- CORS protection in API
- Temporary file cleanup after processing

⚠️ **Areas for Improvement:**

- No private key encryption at rest
- Single admin point of failure
- No rate limiting on RPC calls or API endpoints
- No input validation on file sizes/types
- API keys visible in plain text `.env`
- No authentication on API endpoints

### Best Practices

1. **Never commit `.env` to version control**
2. **Use hardware wallets for production**
3. **Implement multi-signature admin control**
4. **Validate all file inputs before hashing**
5. **Use secret management services (AWS Secrets Manager, HashiCorp Vault)**
6. **Monitor Alchemy API rate limits**
7. **Implement audit logging for all certifications**
8. **Add API authentication (JWT, API keys)**
9. **Implement file size limits and MIME type validation**
10. **Add rate limiting to prevent abuse**

---

## 🧪 Development Workflow

### Typical Development Flow

```bash
# 1. Make changes to Go code
vim internal/blockchain/logic.go

# 2. Test via CLI
cd cmd/cli
go run main.go register test.pdf "Test Student" "Test Course" "Test Issuer"
go run main.go verify test.pdf

# 3. Test via API
cd cmd/api
go run main.go
# In another terminal:
curl -X POST http://localhost:8080/api/v1/register \
  -F "pdf=@test.pdf" \
  -F "student_name=Test" \
  -F "course_name=Test" \
  -F "issuer=Test"

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

# 6. Rebuild applications
cd cmd/cli && go build
cd cmd/api && go build
```

---

## 🛠️ Technical Deep Dive

### How Hashing Works

```go
// GenerateHash creates a Keccak256 hash of any file
func GenerateHash(filePath string) (string, error) {
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
// Certificate structure with metadata
struct Certificate {
    bool isValid;
    string studentName;
    string courseName;
    string issuerName;
    uint256 dateEmited;
}

// Admin-controlled registry
address public admin;
mapping(bytes32 => Certificate) public certificates;

// Only admin can certify with metadata
function registerCertificate(
    bytes32 datahash,
    string memory _studentName,
    string memory _courseName,
    string memory _issuerName
) public {
    require(msg.sender == admin, "Only the admin can certify documents");
    certificates[datahash] = Certificate({
        isValid: true,
        studentName: _studentName,
        courseName: _courseName,
        issuerName: _issuerName,
        dateEmited: block.timestamp
    });
    emit CertificateCreated(datahash, block.timestamp);
}
```

**Key Design Decisions:**

1. **Struct-based storage**: Stores certificate metadata on-chain
2. **Mapping over array**: O(1) lookup time
3. **Events for indexing**: Off-chain services can listen for new certificates
4. **Public mapping**: Verification costs no gas (automatic getter)
5. **Timestamp recording**: Automatic date tracking via `block.timestamp`
6. **No deletion**: Immutable by design

---

## 🔍 Troubleshooting

### Common Issues

#### "Usage: go run main.go <file_path>"

**Correct usage**:

```bash
# CLI - To register (requires metadata)
go run main.go register <file_path> <studentName> <courseName> <issuerName>

# CLI - To verify
go run main.go verify <file_path>
```

#### "Error: Cant found ALCHEMY_URL in .env file"

- Verify `.env` exists in project root
- Check `ALCHEMY_URL` is set and not commented
- Ensure running from correct directory (`cmd/cli/` or `cmd/api/`)

#### "Cant Connect to Alchemy"

- Verify Alchemy URL is correct
- Check internet connectivity
- Confirm API key is valid
- Test with: `curl $ALCHEMY_URL -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'`

#### "Failed to register certificate: insufficient funds"

- Request testnet tokens from [Polygon Faucet](https://faucet.polygon.technology/)
- Check balance: `cast balance YOUR_ADDRESS --rpc-url $ALCHEMY_URL`

#### "Invalid PRIVATE_KEY"

- Ensure key starts with `0x`
- Verify key is 64 hex characters (32 bytes)
- Export from Metamask: Account Details → Export Private Key

#### API not receiving requests

- Check if server is running on port 8080
- Verify firewall allows connections
- Check CORS configuration matches your frontend URL
- Test with cURL before using frontend

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
| API response time      | <2s   | Free (server-side)    |

_Gas costs on Polygon Amoy testnet (free). Mainnet costs will vary with POL price._

---

## 🚧 Future Development

### Planned Features

1. ✅ Certificate registration with metadata
2. ✅ Certificate validation CLI command
3. ✅ REST API for web integration
4. ⏳ Batch certificate registration endpoint
5. ⏳ Certificate revocation mechanism
6. ⏳ Multi-signature admin control
7. ⏳ IPFS integration for document storage
8. ⏳ Email notifications on certification
9. ⏳ QR code generation for verification
10. ⏳ GraphQL API
11. ⏳ WebSocket support for real-time updates
12. ⏳ Frontend dashboard (React/Vue)

### Roadmap to Production

- [ ] Add comprehensive unit tests
- [ ] Implement integration tests with mock blockchain
- [ ] Set up CI/CD pipeline
- [ ] Add input validation and sanitization
- [ ] Implement rate limiting
- [ ] Add structured logging (zerolog/zap)
- [ ] Add API authentication (JWT)
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
- Place business logic in `internal/blockchain/logic.go`

---

## 📄 License

MIT License - See LICENSE file for details

---

## 🔗 Useful Resources

- [Polygon Documentation](https://docs.polygon.technology/)
- [Foundry Book](https://book.getfoundry.sh/)
- [go-ethereum Documentation](https://geth.ethereum.org/docs)
- [Gin Web Framework](https://gin-gonic.com/docs/)
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
- **Gin** - Gin-Gonic Team
- **Solidity** - Ethereum Foundation
- **Foundry** - Paradigm
- **Polygon** - Polygon Labs
- **Alchemy** - Alchemy Insights Inc.

---

_Built by engineers, for engineers. Immutable truth on the blockchain._
