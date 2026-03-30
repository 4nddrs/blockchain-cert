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
| **Database**               | PostgreSQL/Supabase  | Certificate metadata & usage tracking   |
| **API Documentation**      | Swagger/OpenAPI 2.0  | Interactive API documentation           |
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
│       ├── main.go           # REST API server (refactored, minimal)
│       └── temp_/            # Temporary upload directory
├── internal/
│   ├── blockchain/
│   │   ├── logic.go          # Core business logic (register/verify)
│   │   ├── certifyer.go      # Auto-generated contract bindings
│   │   └── client.go         # Ethereum client connection (deprecated)
│   ├── config/
│   │   └── config.go         # Centralized configuration management
│   ├── handlers/
│   │   ├── certificate.go    # HTTP handlers for certificates
│   │   └── admin.go          # HTTP handlers for admin operations
│   ├── middleware/
│   │   └── auth.go           # Authentication middleware
│   ├── repository/
│   │   └── *.go              # Database access layer
│   └── crypto/
│       └── hash.go           # SHA256 file hashing (deprecated)
├── database/
│   └── database.go           # Database connection and initialization
├── models/
│   ├── certificate.go        # Certificate data model
│   ├── institution.go        # Institution data model
│   ├── trial.go              # Trial usage tracking model
│   └── billing_log.go        # Billing/usage log model
├── contracts/
│   └── Certifyer.sol         # Solidity smart contract
├── docs/
│   ├── docs.go               # Swagger documentation (auto-generated)
│   ├── swagger.json          # OpenAPI specification (auto-generated)
│   └── swagger.yaml          # OpenAPI YAML format (auto-generated)
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
- **`cmd/api/`**: REST API server (thin wrapper around handlers)
- **`internal/blockchain/`**: Core business logic and smart contract bindings
- **`internal/config/`**: Configuration management with environment variable validation
- **`internal/handlers/`**: HTTP request handlers separated by domain (certificates, admin)
- **`internal/middleware/`**: Reusable middleware (authentication, logging, etc.)
- **`internal/repository/`**: Database operations and data access layer
- **`database/`**: Database connection management
- **`models/`**: Data structures for database entities
- **`contracts/`**: Solidity smart contracts source code
- **`docs/`**: Auto-generated Swagger/OpenAPI documentation
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
# Blockchain Configuration
ALCHEMY_URL=https://polygon-amoy.g.alchemy.com/v2/YOUR_API_KEY
CONTRACT_ADDRESS=0xYOUR_DEPLOYED_CONTRACT_ADDRESS
PRIVATE_KEY=0xYOUR_PRIVATE_KEY_HERE

# Database Configuration (Supabase)
DATABASE_URL=postgresql://user:password@host:port/database

# Admin Configuration
ADMIN_SECRET=your-secure-admin-secret-token

# Server Configuration (Optional)
SERVER_PORT=8080
CORS_ORIGINS=http://localhost:3000
TEMP_UPLOAD_DIR=./temp_uploads
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

The API has been refactored with clean architecture principles:
- **Separated concerns**: Handlers, middleware, and config are in separate packages
- **Database integration**: Certificate metadata stored in PostgreSQL/Supabase
- **Trial & Institution modes**: Supports both free trial and paid institution accounts
- **Swagger documentation**: Interactive API testing UI included

### Start the API Server

```bash
cd cmd/api
go run main.go

# Server will start on http://localhost:8080
# Swagger UI: http://localhost:8080/swagger/index.html
```

### 🎯 API Features

#### **Dual Usage Modes**

1. **Trial Mode** (No API Key)
   - Free tier for testing
   - Limited registrations per IP/fingerprint
   - No authentication required

2. **Institution Mode** (With API Key)
   - Requires valid `X-API-Key` header
   - Credit-based system
   - Managed via admin endpoints

---

### 📚 Interactive API Documentation (Swagger)

The API includes **interactive Swagger documentation** for easy testing and integration.

#### Access Swagger UI

Once the API server is running, visit:

```
http://localhost:8080/swagger/index.html
```

**Features:**
- ✅ **Try it out** - Test endpoints directly from your browser
- ✅ **Request/Response examples** - See expected formats
- ✅ **Schema definitions** - Understand data structures
- ✅ **Authentication setup** - Configure API keys and admin auth
- ✅ **File upload testing** - Upload PDFs for registration/verification

#### Quick Testing Workflow in Swagger

1. **Register a Certificate (Trial Mode)**
   - Navigate to `POST /api/v1/register`
   - Click "Try it out"
   - Upload a PDF file
   - Fill in student_name, course_name, issuer
   - Click "Execute"
   - Copy the `hash` from response

2. **Verify Certificate**
   - Navigate to `POST /api/v1/verify`
   - Upload the same PDF
   - Click "Execute"
   - See certificate details

3. **Get Certificate by Hash**
   - Navigate to `GET /api/v1/certificates/{hash}`
   - Paste the hash from step 1
   - Click "Execute"

4. **Admin Operations** (Requires Authentication)
   - Click "Authorize" button (top-right)
   - Enter your `ADMIN_SECRET` value
   - Click "Authorize"
   - Now you can access admin endpoints

---

### API Endpoints

#### 🔓 Public Endpoints (No Authentication Required)

#### 1. Register Certificate

**Endpoint**: `POST /api/v1/register`

**Content-Type**: `multipart/form-data`

**Headers** (Optional):
- `X-API-Key`: Institution API key (if not provided, uses trial mode)

**Parameters**:

- `pdf` (file): PDF document to certify
- `student_name` (string): Student's full name
- `course_name` (string): Course name
- `issuer` (string): Issuing organization
- `fingerprint` (string, optional): Browser fingerprint for trial tracking

**Trial Mode Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -F "pdf=@certificate.pdf" \
  -F "student_name=Andres Menchaca" \
  -F "course_name=Blockchain Development" \
  -F "issuer=Tech Academy"
```

**Institution Mode Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "X-API-Key: your-institution-api-key" \
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

**Error Responses**:

- **400 Bad Request**: Missing required fields or invalid data
  ```json
  {
    "error": "student_name, course_name, and issuer are required"
  }
  ```

- **401 Unauthorized**: Invalid API key (institution mode)
  ```json
  {
    "error": "Invalid API key"
  }
  ```

- **402 Payment Required**: Institution out of credits
  ```json
  {
    "error": "Insufficient credits. Please contact support."
  }
  ```

- **409 Conflict**: Certificate already registered
  ```json
  {
    "error": "Certificate already registered",
    "tx_hash": "0x..."
  }
  ```

- **429 Too Many Requests**: Trial limit reached
  ```json
  {
    "error": "Free trial limit reached"
  }
  ```

---

#### 2. Verify Certificate (by File)

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

#### 3. Get Certificate by Hash

**Endpoint**: `GET /api/v1/certificates/:hash`

**Description**: Verify a certificate using its hash directly, without needing the original file.

**Parameters**:

- `hash` (URL parameter): Document hash in hexadecimal format (66 characters, starts with `0x`)

**Example with cURL**:

```bash
curl -X GET http://localhost:8080/api/v1/certificates/0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
```

**Example with Browser**:

```
http://localhost:8080/api/v1/certificates/0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab
```

**Success Response (Valid Certificate)** (200 OK):

```json
{
  "status": "valid",
  "data": {
    "hash": "0xabcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab",
    "student_name": "Andres Menchaca",
    "course_name": "Blockchain Development",
    "issuer": "Tech Academy",
    "date": "2026-03-24 14:30:45"
  }
}
```

**Not Found Response** (404 Not Found):

```json
{
  "status": "not_found",
  "message": "Certificate not found or invalid"
}
```

**Error Response (Invalid Hash Format)** (400 Bad Request):

```json
{
  "status": "error",
  "message": "Invalid hash format"
}
```

**Error Response (Server Error)** (500 Internal Server Error):

```json
{
  "status": "error",
  "message": "Failed to retrieve certificate"
}
```

**Usage Notes**:

- The hash must be exactly 66 characters long (including the `0x` prefix)
- This endpoint is useful for:
  - Verifying certificates from QR codes containing only the hash
  - Building verification links that don't require file uploads
  - Creating public verification pages with shareable URLs
  - Mobile applications that store only hashes

---

#### 🔒 Admin Endpoints (Authentication Required)

Admin endpoints require the `Authorization` header with your admin secret token.

#### 4. Create Institution

**Endpoint**: `POST /api/v1/admin/institutions`

**Headers**:
- `Authorization`: Your admin secret (set in `.env` as `ADMIN_SECRET`)

**Content-Type**: `application/json`

**Request Body**:

```json
{
  "name": "Tech University",
  "email": "contact@techuni.edu",
  "plan": "premium"
}
```

**Plan Options**: `basic`, `premium`, `enterprise`

**Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/admin/institutions \
  -H "Authorization: your-admin-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Tech University",
    "email": "contact@techuni.edu",
    "plan": "premium"
  }'
```

**Success Response** (201 Created):

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Tech University",
  "email": "contact@techuni.edu",
  "plan_type": "premium",
  "api_key": "inst_abc123def456ghi789",
  "credits_remaining": 1000,
  "created_at": "2026-03-27T20:00:00Z"
}
```

**Error Responses**:

- **400 Bad Request**: Invalid input
  ```json
  {
    "error": "Key: 'CreateInstitutionRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
  }
  ```

- **403 Forbidden**: Invalid or missing admin secret
  ```json
  {
    "error": "Unauthorized access"
  }
  ```

---

#### 5. List Institutions

**Endpoint**: `GET /api/v1/admin/institutions`

**Headers**:
- `Authorization`: Your admin secret

**Example with cURL**:

```bash
curl -X GET http://localhost:8080/api/v1/admin/institutions \
  -H "Authorization: your-admin-secret"
```

**Success Response** (200 OK):

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Tech University",
    "email": "contact@techuni.edu",
    "plan_type": "premium",
    "api_key": "inst_abc123def456ghi789",
    "credits_remaining": 950,
    "created_at": "2026-03-27T20:00:00Z"
  },
  {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "name": "Business School",
    "email": "admin@bizschool.com",
    "plan_type": "enterprise",
    "api_key": "inst_xyz789abc456def123",
    "credits_remaining": 5000,
    "created_at": "2026-03-26T15:30:00Z"
  }
]
```

---

#### 6. Add Credits to Institution

**Endpoint**: `POST /api/v1/admin/institutions/{id}/credits`

**Headers**:
- `Authorization`: Your admin secret

**Content-Type**: `application/json`

**URL Parameters**:
- `id`: Institution UUID

**Request Body**:

```json
{
  "additional_credits": 500
}
```

**Example with cURL**:

```bash
curl -X POST http://localhost:8080/api/v1/admin/institutions/550e8400-e29b-41d4-a716-446655440000/credits \
  -H "Authorization: your-admin-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "additional_credits": 500
  }'
```

**Success Response** (200 OK):

```json
{
  "message": "Credits added successfully",
  "institution_id": "550e8400-e29b-41d4-a716-446655440000",
  "credits_added": 500
}
```

**Error Responses**:

- **400 Bad Request**: Invalid input (negative or zero credits)
  ```json
  {
    "error": "Key: 'AddCreditsRequest.AdditionalCredits' Error:Field validation for 'AdditionalCredits' failed on the 'min' tag"
  }
  ```

- **500 Internal Server Error**: Database error
  ```json
  {
    "error": "Failed to add credits"
  }
  ```

---

#### 7. Update Institution Plan

**Endpoint**: `PUT /api/v1/admin/institutions/{id}/plan`

**Headers**:
- `Authorization`: Your admin secret

**Content-Type**: `application/json`

**URL Parameters**:
- `id`: Institution UUID

**Request Body**:

```json
{
  "new_plan": "enterprise"
}
```

**Plan Options**: `basic`, `premium`, `enterprise`

**Example with cURL**:

```bash
curl -X PUT http://localhost:8080/api/v1/admin/institutions/550e8400-e29b-41d4-a716-446655440000/plan \
  -H "Authorization: your-admin-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "new_plan": "enterprise"
  }'
```

**Success Response** (200 OK):

```json
{
  "message": "Plan updated successfully",
  "institution_id": "550e8400-e29b-41d4-a716-446655440000",
  "new_plan": "enterprise"
}
```

**Error Responses**:

- **400 Bad Request**: Invalid plan type
  ```json
  {
    "error": "Key: 'UpdatePlanRequest.NewPlan' Error:Field validation for 'NewPlan' failed on the 'required' tag"
  }
  ```

- **500 Internal Server Error**: Database error
  ```json
  {
    "error": "Failed to update plan"
  }
  ```

---

### Database Integration

The system now stores certificate metadata in a database (PostgreSQL/Supabase) alongside blockchain registration:

**Stored Data:**
- Certificate hash and metadata (student name, course, issuer)
- Transaction hash for blockchain verification
- Institution usage tracking (credits consumed)
- Trial usage tracking (IP address, fingerprint)
- Billing logs for audit trail

**Benefits:**
- ✅ Fast querying without blockchain calls
- ✅ Usage analytics and reporting
- ✅ Credit management for institutions
- ✅ Trial abuse prevention
- ✅ Audit trail for compliance

**Database Schema:**

```
certificates:
  - file_hash (unique)
  - student_name
  - course_name
  - tx_hash
  - created_at

institutions:
  - id (UUID)
  - name
  - email
  - api_key (unique)
  - plan_type
  - credits_remaining
  - created_at

trial_usage:
  - ip_address
  - fingerprint
  - usage_count
  - last_used_at

billing_logs:
  - institution_id
  - action_type
  - cost
  - timestamp
```

---

### CORS Configuration

The API CORS settings are managed centrally via environment variables.

**Default Configuration:**
- Allowed Origins: `http://localhost:3000` (or value from `CORS_ORIGINS` env var)
- Allowed Methods: `GET`, `POST`, `OPTIONS`
- Allowed Headers: `Origin`, `Content-Type`, `X-API-Key`
- Credentials: Enabled
- Max Age: 12 hours

To modify allowed origins, update your `.env`:

```env
CORS_ORIGINS=http://localhost:3000,https://yourdomain.com
```

Or modify `internal/config/config.go` for more complex CORS rules.

---

### 📚 API Documentation (Swagger/OpenAPI)

The API includes **interactive Swagger documentation** for easy testing and integration.

#### Access Swagger UI

Once the API server is running:

```
http://localhost:8080/swagger/index.html
```

This provides:
- **Interactive API testing** - Try endpoints directly from the browser
- **Request/response examples** - See expected formats
- **Schema definitions** - Understand data structures
- **Authentication info** - API security details

#### OpenAPI Specification

The raw OpenAPI 3.0 specification is available at:

```
http://localhost:8080/api/v1/openapi.json
```

Use this JSON file to:
- Generate client SDKs (using tools like OpenAPI Generator)
- Import into Postman/Insomnia
- Integrate with API gateways
- Auto-generate documentation for other platforms

#### Regenerate Swagger Docs

After modifying API endpoints or annotations:

```bash
# Install swag CLI (if not already installed)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate updated documentation from project root
swag init -g cmd/api/main.go --parseDependency --parseInternal

# Restart API server to see changes
cd cmd/api
go run main.go
```

**Swagger Annotations Guide:**

The API uses Go doc comments to generate documentation:

```go
// @Summary Short description
// @Description Detailed explanation
// @Tags category
// @Accept multipart/form-data or json
// @Produce json
// @Param name formData file true "Description"
// @Param X-API-Key header string false "Institution API key"
// @Success 200 {object} ResponseType
// @Failure 400 {object} ErrorType
// @Router /endpoint [post]
```

**Structured Response Types:**

Response models are defined in `internal/handlers/` for better documentation:

```go
// RegisterResponse represents the response for certificate registration
type RegisterResponse struct {
    Status string `json:"status" example:"success"`
    Hash   string `json:"hash" example:"0xabcd1234..."`
    TxHash string `json:"tx_hash" example:"0x1234abcd..."`
}
```

---

### Verify Transaction on Blockchain Explorer

1. Copy the transaction hash from registration
2. Visit [Polygon Amoy Explorer](https://amoy.polygonscan.com/)
3. Paste the transaction hash
4. View the `CertificateCreated` event

---

## 🏛️ Clean Code Architecture

The project follows **clean architecture** and **separation of concerns** principles:

### Layer Structure

```
┌─────────────────────────────────────────────────┐
│          Presentation Layer (cmd/)              │
│      CLI Interface    |    API Server           │
└──────────────┬──────────────────────────────────┘
               │
┌──────────────▼──────────────────────────────────┐
│         Handler Layer (internal/handlers/)      │
│  Certificate Handlers  |  Admin Handlers        │
└──────────────┬──────────────────────────────────┘
               │
┌──────────────▼──────────────────────────────────┐
│      Business Logic (internal/blockchain/)      │
│  GenerateHash | RegisterCertificate | Verify    │
└──────────────┬──────────────────────────────────┘
               │
┌──────────────▼──────────────────────────────────┐
│    Data Access Layer (internal/repository/)     │
│  Database Queries  |  Transaction Management    │
└──────────────┬──────────────────────────────────┘
               │
┌──────────────▼──────────────────────────────────┐
│       Infrastructure (database/, models/)       │
│   PostgreSQL   |   Ethereum Client   |  Models  │
└─────────────────────────────────────────────────┘
```

### Core Business Logic (`internal/blockchain/logic.go`)

All certificate operations centralized:

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

### Handler Layer (`internal/handlers/`)

HTTP request handling separated by domain:

- **`certificate.go`**: Certificate registration and verification handlers
- **`admin.go`**: Institution management handlers

Each handler:
- ✅ Validates input
- ✅ Calls business logic
- ✅ Returns structured responses
- ✅ Handles errors appropriately

### Repository Layer (`internal/repository/`)

Database operations abstracted:

- Certificate CRUD operations
- Institution management
- Trial usage tracking
- Billing log recording

### Configuration (`internal/config/`)

Centralized configuration management:

- Environment variable loading
- Validation
- Default values
- Error handling

### Middleware (`internal/middleware/`)

Reusable cross-cutting concerns:

- Admin authentication
- (Future: Rate limiting, logging, metrics)

### Benefits of Current Architecture

✅ **Single Responsibility**: Each module has one clear purpose  
✅ **DRY Principle**: Logic shared between interfaces  
✅ **Testability**: Pure functions can be unit tested  
✅ **Maintainability**: Changes isolated to specific layers  
✅ **Scalability**: Easy to add new interfaces  
✅ **Separation of Concerns**: UI, business logic, and data access decoupled

### Interface Layer

Both CLI and API are thin wrappers:

- **CLI** (`cmd/cli/main.go`): Parses arguments → calls logic → formats output
- **API** (`cmd/api/main.go`): Routes requests → initializes dependencies → delegates to handlers

This makes it trivial to add new interfaces (gRPC, WebSocket, GraphQL) without duplicating business logic.

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
