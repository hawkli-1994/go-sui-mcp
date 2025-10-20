# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go Sui MCP is a Management Control Plane (MCP) server that wraps the Sui blockchain CLI client, exposing Sui operations as MCP tools. It supports both stdio and SSE (Server-Sent Events) transport modes for integration with IDEs like Cursor.

## Build and Run Commands

```bash
# Build the binary
make build
# or
go build -o go-sui-mcp main.go

# Run in stdio mode (default)
./go-sui-mcp server

# Run in SSE mode
./go-sui-mcp server --sse --port 8080

# Run directly with go
make run
```

## Architecture

The codebase follows a three-layer architecture:

### 1. Client Layer (`internal/sui/client.go`)
- **`Client` struct**: Wraps the Sui CLI executable
- **`ExecuteCommand(args ...string)`**: Low-level command execution that shells out to the `sui` binary
- Uses `viper.GetString("sui.executable_path")` to locate the Sui binary (defaults to "sui")
- Returns raw stdout from Sui CLI commands

### 2. Service Layer (`internal/services/sui_service.go`)
- **`SuiService` struct**: Provides business logic on top of the Client
- Methods follow MCP handler signature: `(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)`
- Responsible for:
  - Extracting and validating parameters from MCP requests
  - Calling the appropriate Client methods
  - Formatting responses as MCP tool results

### 3. Tools Layer (`internal/services/sui_tools.go`)
- **`SuiTools` struct**: Defines MCP tool schemas using the `mcp-go` library
- Each method returns a `mcp.Tool` with:
  - Tool name (e.g., "sui-balance-summary")
  - Parameter schemas (string, number, required/optional)
  - Descriptions for the MCP protocol

### Handler Registration (`cmd/server.go`)
The `registerHandlers()` function connects tools to their service implementations:
```go
s.AddTool(suiTools.GetBalanceSummary(), suiService.GetBalanceSummary)
```

## Available MCP Tools

The server exposes 28 Sui operations as MCP tools, organized by category:

### Version and Path (2 tools)
1. **sui-formatted-version**: Get Sui client version
2. **sui-path**: Get path to the local sui binary

### Address and Environment Management (5 tools)
3. **sui-active-address**: Get the current active address
4. **sui-addresses**: List all addresses managed by the client
5. **sui-active-env**: Get current active environment (devnet/testnet/mainnet)
6. **sui-envs**: List all configured network environments
7. **sui-chain-identifier**: Query chain identifier from RPC endpoint

### Balance and Objects (3 tools)
8. **sui-balance-summary**: Get balance for an address (optional address param)
9. **sui-objects-summary**: Get objects owned by an address (returns JSON)
10. **sui-object**: Get details of a specific object (required objectID)

### Gas Management (2 tools)
11. **sui-gas**: Get all gas objects owned by an address
12. **sui-faucet**: Request test coins from faucet (devnet/testnet only)

### Transaction Operations (7 tools)
13. **sui-transfer**: Transfer an object to another address
14. **sui-transfer-sui**: Simplified SUI transfer (with optional amount)
15. **sui-pay-sui**: Pay SUI to recipients (original implementation)
16. **sui-pay**: Pay coins to multiple recipients with specified amounts
17. **sui-pay-all-sui**: Pay all residual SUI after deducting gas
18. **sui-split-coin**: Split a coin object into multiple coins
19. **sui-merge-coin**: Merge two coin objects into one

### Transaction Info (1 tool)
20. **sui-process-transaction**: Get transaction details (required txID)

### Contract Interaction (3 tools)
21. **sui-call**: Call a Move function on the blockchain
22. **sui-publish**: Publish Move modules to the blockchain
23. **sui-dynamic-field**: Query a dynamic field by parent object ID

### Move Development (3 tools)
24. **sui-move-build**: Build a Move package
25. **sui-move-test**: Run Move unit tests
26. **sui-move-new**: Create a new Move package

### Keytool Management (3 tools)
27. **sui-keytool-list**: List all keys in the keystore
28. **sui-keytool-generate**: Generate a new keypair (ed25519/secp256k1/secp256r1)
29. **sui-keytool-export**: Export private key in Bech32 format

## Critical Context from .cursorrules

**IMPORTANT**: When working with Sui-related queries:
- Always check if an MCP tool is available before implementing custom solutions
- Prefer using MCP tools over direct implementation
- Do NOT cache MCP tool results - call the tools each time to get fresh data
- This ensures consistency with live blockchain state

## Configuration

Uses Viper for hierarchical configuration (priority order):
1. Command-line flags (`--port`, `--sse`)
2. Environment variables (prefix: `GOSUI_`, e.g., `GOSUI_SERVER_PORT`)
3. Config file (`$HOME/.go-sui-mcp.yaml` or via `--config` flag)

Example config file:
```yaml
server:
  port: 8080
  sse: false
sui:
  executable_path: "sui"
```

## Server Modes

### stdio Mode (Default)
- Communicates via standard input/output
- Used for direct process spawning by IDEs
- Start with: `./go-sui-mcp server`

### SSE Mode
- HTTP server with Server-Sent Events
- Useful for development and remote connections
- Start with: `./go-sui-mcp server --sse --port 8080`

## Dependencies

- **github.com/mark3labs/mcp-go**: MCP protocol implementation
- **github.com/spf13/cobra**: CLI framework
- **github.com/spf13/viper**: Configuration management
- **Sui CLI**: Must be installed and in PATH (or configured via `sui.executable_path`)

## Testing

No test files currently exist in the codebase. When adding tests, use Go's standard testing package with the `_test.go` suffix.
