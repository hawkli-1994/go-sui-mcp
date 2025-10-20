# Go Sui MCP (Management Control Plane)
![0e0d353d18274ae38d5e695991b99fba_1](https://github.com/user-attachments/assets/faf47f7f-2cef-46a1-90b5-5d8423f8a8f2)


A Go-based management control plane server for Sui blockchain, providing MCP (Management Control Plane) tools to interact with local Sui client commands. This project integrates with Cursor IDE for enhanced development experience.

## Features

- MCP tools for Sui client operations
- Support for both stdio and SSE (Server-Sent Events) modes
- Integration with Cursor IDE
- Configuration via config file, environment variables, or command-line flags

## Prerequisites

- Go 1.20 or higher
- Sui client installed and available in PATH
- Cursor IDE (for development)

## Installation

```bash
# Clone the repository
git clone https://github.com/krli/go-sui-mcp.git
cd go-sui-mcp

# Build the application
go build -o go-sui-mcp
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

### Example Tool Usage (in Cursor)

Transfer SUI tokens:
```typescript
await mcp.invoke("sui-pay-sui", {
  recipient: "0x...",
  amounts: 1000000000, // 1 SUI
  "gas-budget": "2000000",
  "input-coins": "0x..."
});
```

Get balance summary:
```typescript
await mcp.invoke("sui-balance-summary", {
  address: "0x..." // optional, uses current address if not provided
});
```

## Development

The project uses the following key components:

- `internal/sui`: Core Sui client implementation
- `internal/services`: MCP tools and service implementations
- `cmd`: Command-line interface implementation

## License

See [LICENSE](LICENSE) file.
