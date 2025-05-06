package services

import (
	"github.com/mark3labs/mcp-go/mcp"
	// "github.com/mark3labs/mcp-go/server"
)

type SuiTools struct {
}

func NewSuiTools() *SuiTools {
	return &SuiTools{}
}

func (s *SuiTools) GetFormattedVersion() mcp.Tool {
	return mcp.NewTool(
		"sui-formatted-version",
		mcp.WithDescription("Get the formatted version of the Sui client"),
	)
}

func (s *SuiTools) GetBalanceSummary() mcp.Tool {
	return mcp.NewTool(
		"sui-balance-summary",
		mcp.WithString("address",
			mcp.Required(),
			mcp.Description("Address to get the balance summary of"),
		),
		mcp.WithDescription("Get the balance summary of the Sui client"),
	)
}

func (s *SuiTools) GetObjectsSummary() mcp.Tool {
	return mcp.NewTool(
		"sui-objects-summary",
		mcp.WithString("address",
			mcp.Required(),
			mcp.Description("Address to get the objects summary of"),
		),
		mcp.WithDescription("Get the objects summary of the Sui client"),
	)
}

func (s *SuiTools) ProcessTransaction() mcp.Tool {
	return mcp.NewTool(
		"sui-process-transaction",
		mcp.WithString("txID",
			mcp.Required(),
			mcp.Description("Transaction ID to process"),
		),
		mcp.WithDescription("Process a transaction"),
	)
}

func (s *SuiTools) TransferTokens() mcp.Tool {
	return mcp.NewTool(
		"sui-transfer-tokens",
		mcp.WithString("recipient",
			mcp.Required(),
			mcp.Description("Recipient address"),
		),
		mcp.WithNumber("amount",
			mcp.Required(),
			mcp.Description("Amount to transfer"),
		),
		mcp.WithString("gasBudget",
			mcp.Required(),
			mcp.Description("Gas budget"),
		),
		mcp.WithDescription("Transfer tokens"),
	)
}
