# Go Sui MCP (Management Control Plane)
![0e0d353d18274ae38d5e695991b99fba_1](https://github.com/user-attachments/assets/faf47f7f-2cef-46a1-90b5-5d8423f8a8f2)


A Go-based management control plane server for Sui blockchain, providing MCP (Management Control Plane) tools to interact with local Sui client commands. This project integrates with Cursor IDE for enhanced development experience.

## Features

- **28 Comprehensive MCP Tools**: Complete coverage of Sui blockchain operations
  - Address and environment management
  - Balance and gas operations
  - Transaction handling (transfer, split, merge coins)
  - Smart contract interaction (call, publish)
  - Move development workflow (build, test, new package)
  - Keystore management
- **Dual Transport Modes**: stdio (default) and SSE (Server-Sent Events)
- **IDE Integration**: Works with Cursor, Claude Code, and any MCP-compatible client
- **Flexible Configuration**: Support for config files, environment variables, and CLI flags
- **Zero Dependencies**: Only requires Sui CLI to be installed locally

## Prerequisites

- Go 1.23 or higher
- Sui client installed and available in PATH
- Cursor IDE or any MCP-compatible client (for using the tools)

## Installation

```bash
# Clone the repository
git clone https://github.com/hawkli-1994/go-sui-mcp.git
cd go-sui-mcp

# Build the application
make build
# or
go build -o go-sui-mcp main.go
```

## Configuration

Configuration can be done via:

1. Config file (default: `$HOME/.go-sui-mcp.yaml`)
2. Environment variables
3. Command-line flags

Example config file:

```yaml
server:
  port: 8080
  sse: false
sui:
  executable_path: "sui"
```

Environment variables:

```bash
GOSUI_SERVER_PORT=8080
GOSUI_SERVER_SSE=false
GOSUI_SUI_EXECUTABLE_PATH=sui
```

## Running the Server

The server can be run in two modes:

1. stdio mode (default):
```bash
./go-sui-mcp server
```

2. SSE mode:
```bash
./go-sui-mcp server --sse --port 8080
```

## Cursor IDE Integration

To integrate with Cursor IDE, create a `.cursor/mcp.json` file in your project root:

```json
{
  "mcpServers": {
    "sui-sse": {
      "command": "/path/to/go-sui-mcp",
      "args": ["server", "--sse"]
    },
    "sui-dev": {
      "url": "http://localhost:8080/sse"
    },
    "sui": {
      "command": "/path/to/go-sui-mcp",
      "args": ["server"]
    }
  }
}
```

## Available MCP Tools

The server provides **28 MCP tools** covering comprehensive Sui blockchain operations:

### Version and Path (2 tools)
- `sui-formatted-version`: Get the formatted version of the Sui client
- `sui-path`: Get the path of the local sui binary

### Address and Environment Management (5 tools)
- `sui-active-address`: Get the current active address
- `sui-addresses`: List all addresses managed by the client
- `sui-active-env`: Get current active environment (devnet/testnet/mainnet)
- `sui-envs`: List all configured network environments
- `sui-chain-identifier`: Query chain identifier from RPC endpoint

### Balance and Objects (3 tools)
- `sui-balance-summary`: Get the balance summary of an address
- `sui-objects-summary`: Get the objects summary of an address
- `sui-object`: Get details of a specific object

### Gas Management (2 tools)
- `sui-gas`: Get all gas objects owned by an address
- `sui-faucet`: Request test coins from faucet (devnet/testnet only)

### Transaction Operations (7 tools)
- `sui-transfer`: Transfer an object to another address
- `sui-transfer-sui`: Simplified SUI transfer (with optional amount)
- `sui-pay-sui`: Pay SUI to recipients
- `sui-pay`: Pay coins to multiple recipients with specified amounts
- `sui-pay-all-sui`: Pay all residual SUI after deducting gas
- `sui-split-coin`: Split a coin object into multiple coins
- `sui-merge-coin`: Merge two coin objects into one

### Transaction Info (1 tool)
- `sui-process-transaction`: Process and get details of a transaction

### Contract Interaction (3 tools)
- `sui-call`: Call a Move function on the blockchain
- `sui-publish`: Publish Move modules to the blockchain
- `sui-dynamic-field`: Query a dynamic field by parent object ID

### Move Development (3 tools)
- `sui-move-build`: Build a Move package
- `sui-move-test`: Run Move unit tests
- `sui-move-new`: Create a new Move package

### Keytool Management (3 tools)
- `sui-keytool-list`: List all keys in the keystore
- `sui-keytool-generate`: Generate a new keypair (ed25519/secp256k1/secp256r1)
- `sui-keytool-export`: Export private key in Bech32 format

## Example Tool Usage

### Get current active address
```typescript
await mcp.invoke("sui-active-address");
```

### Get balance summary
```typescript
await mcp.invoke("sui-balance-summary", {
  address: "0x..." // optional, uses current address if not provided
});
```

### Transfer SUI tokens
```typescript
await mcp.invoke("sui-pay-sui", {
  recipient: "0x...",
  amounts: 1000000000, // 1 SUI in MIST
  "gas-budget": "2000000",
  "input-coins": "0x..."
});
```

### Request test tokens from faucet (devnet/testnet)
```typescript
await mcp.invoke("sui-faucet");
```

### Call a Move function
```typescript
await mcp.invoke("sui-call", {
  package: "0x...",
  module: "module_name",
  function: "function_name",
  "type-args": ["0x2::sui::SUI"],
  args: ["arg1", "arg2"],
  "gas-budget": "10000000"
});
```

### Build a Move package
```typescript
await mcp.invoke("sui-move-build", {
  "package-path": "./my_move_package"
});
```

### Run Move tests
```typescript
await mcp.invoke("sui-move-test", {
  "package-path": "./my_move_package",
  filter: "test_name" // optional
});
```

## Project Structure

The project follows a clean three-layer architecture:

```
go-sui-mcp/
├── cmd/                      # CLI commands
│   ├── root.go              # Root command and config initialization
│   └── server.go            # MCP server command and tool registration
├── internal/
│   ├── sui/                 # Sui client layer
│   │   └── client.go        # Wraps Sui CLI commands
│   ├── services/            # Service layer
│   │   ├── sui_service.go   # MCP request handlers
│   │   └── sui_tools.go     # MCP tool definitions
│   └── config/              # Configuration management
│       └── config.go
├── main.go                  # Application entry point
├── Makefile                 # Build automation
├── go.mod                   # Go module dependencies
└── CLAUDE.md               # Claude Code integration guide
```

### Key Components

- **Client Layer** (`internal/sui/client.go`): Wraps Sui CLI commands and executes them
- **Service Layer** (`internal/services/sui_service.go`): Implements MCP handlers with parameter validation
- **Tools Layer** (`internal/services/sui_tools.go`): Defines MCP tool schemas and descriptions
- **Command Layer** (`cmd/`): Cobra-based CLI with server initialization and tool registration

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (when available): `go test ./...`
5. Build and test: `make build && ./go-sui-mcp server`
6. Commit your changes (`git commit -m 'Add some amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

See [LICENSE](LICENSE) file.

## Acknowledgments

- Built with [mcp-go](https://github.com/mark3labs/mcp-go) - Go implementation of Model Context Protocol
- Powered by [Sui](https://sui.io/) blockchain
