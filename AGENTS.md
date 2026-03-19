# AGENTS.md

## Project Overview

**blockchain-cert** is a Go-based blockchain certificate verification system that integrates with Ethereum (via Alchemy) to certify and validate document hashes on-chain.

### Tech Stack
- **Language**: Go 1.25.5
- **Blockchain**: Ethereum via go-ethereum client library
- **Network**: Polygon Amoy testnet (via Alchemy)
- **Smart Contract**: Solidity 0.8.19
- **Key Dependencies**:
  - `github.com/ethereum/go-ethereum v1.17.1` - Ethereum client
  - `github.com/joho/godotenv v1.5.1` - Environment variable management

### Project Purpose
Generates SHA256 hashes of documents (PDFs) and interacts with an Ethereum smart contract to register and validate certificates on the blockchain. The admin can certify documents, and anyone can verify if a document has been certified by checking its hash.

---

## Project Structure

```
blockchain-cert/
├── cmd/
│   └── cli/
│       ├── main.go           # CLI entry point
│       └── title.pdf          # Example test file
├── internal/
│   ├── blockchain/
│   │   └── client.go          # Ethereum client connection logic
│   └── crypto/
│       └── hash.go            # SHA256 file hashing
├── contracts/
│   └── Certificador.sol       # Solidity smart contract
├── .env                       # Environment variables (not committed)
├── go.mod                     # Go module definition
└── go.sum                     # Go dependencies lockfile
```

### Key Directories
- **`cmd/cli/`**: Command-line interface for running the application
- **`internal/blockchain/`**: Ethereum client connectivity
- **`internal/crypto/`**: Cryptographic utilities (file hashing)
- **`contracts/`**: Solidity smart contracts

---

## Essential Commands

### Build & Run
```bash
# Run the CLI (from project root)
cd cmd/cli
go run main.go <file_path>

# Example:
go run main.go title.pdf

# Build the CLI
go build -o blockchain-cert ./cmd/cli
```

### Dependency Management
```bash
# Download dependencies
go mod download

# Update dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### Testing
Currently **no test files exist** in the codebase. When adding tests:
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in specific package
go test ./internal/crypto
```

---

## Code Organization & Patterns

### Package Structure
- **`internal/`** packages are project-private (cannot be imported by external projects)
- Standard Go project layout with `cmd/` for executables and `internal/` for private code
- Each package has a single responsibility (crypto, blockchain)

### Import Paths
Use full module path: `github.com/4nddrs/blockchain-cert/internal/...`

### Error Handling Pattern
```go
// Check error immediately after operation
result, err := SomeFunction()
if err != nil {
    return err  // or log.Fatal for CLI
}
```

Current code uses `log.Fatal()` or `log.Fatalf()` in main.go for unrecoverable errors.

### Context Usage
The blockchain client uses context with timeout:
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```
**Pattern**: 10-second timeout for Ethereum RPC calls

---

## Environment Variables

The project uses `.env` file loaded via `godotenv`. Required variables:

| Variable | Purpose | Example |
|----------|---------|---------|
| `ALCHEMY_URL` | Polygon Amoy RPC endpoint | `https://polygon-amoy.g.alchemy.com/v2/...` |
| `ALCHEMY_API_KEY` | Alchemy API key | (hex string) |
| `CONTRACT_ADDRESS` | Deployed Certificador contract address | `0x1234...` |

### Loading Pattern
```go
err := godotenv.Load("../../.env")  // Relative path from cmd/cli/
if err != nil {
    log.Fatalf("Error loading .env file: %v", err)
}
url := os.Getenv("ALCHEMY_URL")
if url == "" {
    log.Fatal("Error: ALCHEMY_URL no encontrada en el .env")
}
```

**Gotcha**: `.env` path is relative to where `main.go` is executed. From `cmd/cli/`, it uses `../../.env`.

---

## Smart Contract (Certificador.sol)

### Contract Structure
```solidity
// Admin-controlled certificate registry
address public admin;                           // Contract deployer
mapping(bytes32 => bool) public certificados;  // hash => is_certified

// Register certificate (admin only)
function registrarCertificado(bytes32 datahash) public

// Validate certificate (anyone can check)
function validarCertificado(bytes32 datahash) public view returns (bool)

// Event emitted on certification
event CertificadoCreado(bytes32 indexed datahash, uint256 timestamp)
```

### Key Design Decisions
1. **Admin-only registration**: Only the contract deployer can certify documents
2. **Public validation**: Anyone can verify if a hash is certified
3. **Hash storage**: Uses `bytes32` for document hashes (matches SHA256 output size)
4. **Immutable records**: Once certified, always certified (no revocation)

### Deployment
**Note**: No deployment scripts exist yet. The contract needs to be deployed to Polygon Amoy testnet and the address added to `.env`.

---

## Cryptographic Hashing

### Hash Format
```go
// internal/crypto/hash.go
func GenerateFileHash(filePath string) (string, error)
```

- **Algorithm**: SHA256
- **Output Format**: Hex string with `0x` prefix (e.g., `0xabcd1234...`)
- **Rationale**: SHA256 is widely used, secure, and matches Solidity `bytes32` type

### Usage Pattern
```go
hash, err := crypto.GenerateFileHash("document.pdf")
if err != nil {
    return err
}
// hash is ready for blockchain submission
```

---

## Blockchain Integration

### Connection Pattern
```go
// internal/blockchain/client.go
client, err := blockchain.Connect(os.Getenv("ALCHEMY_URL"))
if err != nil {
    log.Fatalf("Cant Connect to Alchemy", err)
}
defer client.Close()
```

### Current State
- ✅ Establishes connection to Polygon Amoy via Alchemy
- ❌ Contract interaction not yet implemented (no ABI bindings)

### Next Steps for Full Integration
1. Generate Go bindings from `Certificador.sol`:
   ```bash
   # Requires solc and abigen
   solc --abi contracts/Certificador.sol -o build
   abigen --abi=build/Certificador.abi --pkg=contracts --out=internal/contracts/certificador.go
   ```
2. Implement `registrarCertificado()` call
3. Implement `validarCertificado()` query
4. Handle transaction signing (requires private key in `.env`)

---

## Code Style & Conventions

### Naming
- **Packages**: lowercase, single word (crypto, blockchain)
- **Files**: lowercase with underscores (client.go, hash.go)
- **Functions**: CamelCase, exported start with capital (GenerateFileHash, Connect)
- **Variables**: camelCase for local, CamelCase for exported

### Comments
- **Exported functions** have doc comments starting with function name
- Example: `// GenerateFileHash take the route of the file and return the sha256 of the file.`
- Note: Some comments are in Spanish (mixed language usage)

### Language Usage
- **Code**: English (function names, variable names)
- **Comments**: Mixed English/Spanish
- **Error messages**: Spanish (`"Error: ALCHEMY_URL no encontrada en el .env"`)

**Recommendation**: Standardize on English for consistency with Go ecosystem.

### Error Messages
Current style uses Spanish. Examples:
- `"Error: ALCHEMY_URL no encontrada en el .env"`
- `"Only the admin can certify documents"` (in contract)

---

## Gotchas & Non-Obvious Patterns

### 1. Working Directory Sensitivity
**Issue**: `.env` loading uses relative path `../../.env` from `cmd/cli/main.go`

**Solution**: Always run from `cmd/cli/` directory or adjust path:
```bash
cd cmd/cli
go run main.go <file>
```

### 2. Unused Client Variable
In `main.go:41`: `_ = client` indicates blockchain client is connected but not yet used for transactions.

### 3. GOTMPDIR Comment
Line 3 in `main.go` has commented export: `// export GOTMPDIR=~/go-cache/tmp`

This suggests the developer may be managing Go's temp directory for build caching or disk space reasons.

### 4. Contract Address Not Used
`.env` has `CONTRACT_ADDRESS` but it's not loaded/used in current code. Will be needed when contract interactions are implemented.

### 5. No Input Validation
`GenerateFileHash()` and main CLI don't validate:
- File existence before hashing
- File type (should it only accept PDFs?)
- Hash format before blockchain submission

### 6. Error Handling Inconsistency
- `main.go:32`: Ignores error from `GenerateFileHash()` (uses `_`)
- `main.go:38`: Uses `log.Fatalf()` but passes `err` as second arg without formatting

### 7. No Graceful Shutdown
Ethereum client is not explicitly closed. Add `defer client.Close()` when client is used.

---

## Testing Strategy

### Current State
**No tests exist.** When adding tests, follow these patterns:

### Recommended Test Structure
```
internal/
├── crypto/
│   ├── hash.go
│   └── hash_test.go        # Test GenerateFileHash
└── blockchain/
    ├── client.go
    └── client_test.go      # Test Connect (may need mocks)
```

### Test Patterns to Follow
```go
// Unit test example
func TestGenerateFileHash(t *testing.T) {
    // Create temp file
    tmpfile, _ := os.CreateTemp("", "test*.pdf")
    defer os.Remove(tmpfile.Name())
    
    tmpfile.Write([]byte("test content"))
    tmpfile.Close()
    
    hash, err := GenerateFileHash(tmpfile.Name())
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    
    if !strings.HasPrefix(hash, "0x") {
        t.Errorf("hash should start with 0x, got %s", hash)
    }
}
```

### Integration Testing
Blockchain tests will need:
- Mock Ethereum client or local testnet (ganache/hardhat)
- Test fixtures for contracts
- Separate test environment variables

---

## Dependencies & External Services

### Alchemy
- **Service**: Ethereum node provider
- **Network**: Polygon Amoy testnet
- **Purpose**: Submit transactions and query blockchain state
- **API Key Required**: Yes (in `.env`)

### Go Ethereum (geth)
- Version: v1.17.1
- Provides: Ethereum client, ABI encoding, transaction signing
- Large dependency tree (~200 lines in go.sum)

---

## Development Workflow

### Typical Development Flow
1. **Make changes** to Go code
2. **Run from CLI**:
   ```bash
   cd cmd/cli
   go run main.go test_document.pdf
   ```
3. **Check connection**: Verify "Success connecting to Alchemy" message
4. **Debug**: Check environment variables and network connectivity

### When Modifying Smart Contract
1. Update `contracts/Certificador.sol`
2. Compile with solc
3. Deploy to Polygon Amoy testnet
4. Update `CONTRACT_ADDRESS` in `.env`
5. Regenerate Go bindings with abigen

### Adding New Features
1. Create package in `internal/` if needed
2. Follow existing error handling patterns
3. Use context with timeout for blockchain calls
4. Add environment variables to `.env` if needed
5. Update this AGENTS.md document

---

## Security Considerations

### Sensitive Data
- **`.env` file**: Contains API keys and should NEVER be committed
- **Private keys**: Not yet implemented but will need secure storage
- **Contract admin**: Single point of failure (only admin can certify)

### Best Practices
1. Never log private keys or API keys
2. Validate all file inputs before hashing
3. Use context timeouts for all network calls (currently 10s)
4. Consider rate limiting for API calls to Alchemy

### Current Security Gaps
- No private key management yet
- No input sanitization on file paths
- No validation of hash format before submission
- API key visible in `.env` (should use secret management)

---

## Future Development Areas

Based on current code state, these areas need implementation:

1. **Contract Interaction**:
   - Generate Go bindings from Solidity ABI
   - Implement transaction signing
   - Add private key to environment variables
   - Call `registrarCertificado()` function
   - Query `validarCertificado()` function

2. **CLI Improvements**:
   - Add `validate` subcommand to check existing certificates
   - Add `register` subcommand to certify new documents
   - Better error messages and help text
   - Input validation

3. **Testing**:
   - Unit tests for crypto package
   - Integration tests with mock blockchain
   - End-to-end tests on testnet

4. **Error Handling**:
   - Return errors instead of log.Fatal in libraries
   - Structured error types
   - Better error context

5. **Configuration**:
   - Support multiple networks (mainnet, testnet)
   - Configuration validation at startup
   - Support for different contract addresses per network

---

## Troubleshooting

### "Error: ALCHEMY_URL no encontrada en el .env"
- Check `.env` file exists in project root
- Verify `ALCHEMY_URL` is set and not commented
- Check working directory when running (should be in `cmd/cli/`)

### "Cant Connect to Alchemy"
- Verify Alchemy URL is correct
- Check internet connectivity
- Verify Alchemy API key is valid
- Check if Polygon Amoy testnet is operational

### "No such file or directory" when hashing
- Provide absolute path or correct relative path
- Verify file exists: `ls -la <file_path>`

### Go module errors
```bash
go mod tidy
go mod download
```

---

## Additional Notes

### Project Maturity
This is an **early-stage project**:
- ✅ Basic structure in place
- ✅ File hashing working
- ✅ Blockchain connection established
- ✅ Smart contract designed
- ❌ No contract deployment scripts
- ❌ No contract interaction code
- ❌ No tests
- ❌ No CI/CD

### Performance Considerations
- SHA256 hashing is fast for typical documents
- Ethereum transactions can take 2-30 seconds depending on network
- Alchemy free tier has rate limits

### Recommended Tools
- **solc**: Solidity compiler (needed for contract compilation)
- **abigen**: Generate Go bindings from ABI (part of go-ethereum)
- **hardhat/foundry**: For local testing of smart contracts
- **metamask**: For manual contract interaction during development

---

*Last updated: Based on codebase analysis 2026-03-19*
