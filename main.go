package main

import (
	"log"

	"github.com/krli/go-sui-mcp/cmd"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

// package main

// import (
//     "context"
//     "errors"
//     "fmt"

//     "github.com/mark3labs/mcp-go/mcp"
//     "github.com/mark3labs/mcp-go/server"
// )

// func main() {
//     // Create MCP server
//     s := server.NewMCPServer(
//         "Demo 🚀",
//         "1.0.0",
//     )

//     // Add tool
//     tool := mcp.NewTool("hello_world",
//         mcp.WithDescription("Say hello to someone"),
//         mcp.WithString("name",
//             mcp.Required(),
//             mcp.Description("Name of the person to greet"),
//         ),
//     )

//     // Add tool handler
//     s.AddTool(tool, helloHandler)

//     // Start the stdio server
//     if err := server.ServeStdio(s); err != nil {
//         fmt.Printf("Server error: %v\n", err)
//     }
// }

// func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
//     name, ok := request.Params.Arguments["name"].(string)
//     if !ok {
//         return nil, errors.New("name must be a string")
//     }

//     return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
// }
