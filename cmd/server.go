package cmd

import (
	"log"

	"github.com/krli/go-sui-mcp/internal/services"
	"github.com/krli/go-sui-mcp/internal/sui"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port int
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the MCP server",
	Long:  `Start the Management Control Plane server to handle Sui client operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Local flags for the server command
	serverCmd.Flags().IntVar(&port, "port", 8080, "Port to run the server on")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
}

func startServer() {
	// Load configuration
	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Fatalf("Failed to load configuration: %v", err)
	// }

	// Create a new Sui client
	suiClient := sui.NewClient()

	// Create service layer
	suiService := services.NewSuiService(suiClient)
	suiTools := services.NewSuiTools()
	s := server.NewMCPServer(
		"SUI MCP",
		"1.0.0",
	)
	s.AddTool(suiTools.GetFormattedVersion(), suiService.GetFormattedVersion)
	s.AddTool(suiTools.GetBalanceSummary(), suiService.GetBalanceSummary)
	s.AddTool(suiTools.GetObjectsSummary(), suiService.GetObjectsSummary)
	s.AddTool(suiTools.ProcessTransaction(), suiService.ProcessTransaction)
	s.AddTool(suiTools.TransferTokens(), suiService.TransferTokens)

	// fmt.Println("Starting server...")
	sseServer := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8080"))
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
	// if err := server.ServeStdio(s); err != nil {

	//     fmt.Printf("Server error: %v\n", err)
	// }

	// // Initialize the Gin router
	// router := gin.Default()

	// // Register routes
	// router.GET("/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "ok",
	// 	})
	// })

	// // Register Sui client API handlers
	// handlers.RegisterSuiHandlers(router, suiClient, suiService)

	// // Start the server
	// addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	// log.Printf("Starting server on %s", addr)
	// if err := router.Run(addr); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}
