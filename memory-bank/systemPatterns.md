# System Patterns: Go-Sui-MCP

## Architecture
- Layered architecture with clear separation of concerns:
  1. CLI Layer (cmd package)
  2. Service Layer (services package)
  3. Client Layer (sui package)
  4. Configuration Layer (config package)
  5. MCP Integration Layer

## Key Technical Decisions
- Cobra for CLI implementation
- Viper for configuration management
- MCP-Go for protocol implementation
- Modular package structure
- Interface-based design
- Configuration via file, env vars, and flags
- Support for both stdio and SSE modes

## Design Patterns
- Command pattern for CLI operations
- Factory pattern for service creation
- Strategy pattern for server modes
- Dependency injection for services
- Singleton pattern for configuration
- Builder pattern for MCP tools
- Observer pattern for SSE events

## Component Relationships
- CLI Layer → Service Layer → Client Layer
- Configuration Layer → All Layers
- Service Layer ↔ MCP Layer
- Client Layer → Sui Network
- MCP Layer → External Tools 