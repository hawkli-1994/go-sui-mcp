package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"context"
	"errors"

	"github.com/krli/go-sui-mcp/internal/sui"
	"github.com/mark3labs/mcp-go/mcp"
)

// SuiService provides higher-level operations on the Sui blockchain
type SuiService struct {
	client *sui.Client
}

// NewSuiService creates a new Sui service
func NewSuiService(client *sui.Client) *SuiService {
	return &SuiService{
		client: client,
	}
}

// GetFormattedVersion returns a cleaned version string
func (s *SuiService) GetFormattedVersion(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	version, err := s.client.GetVersion()
	if err != nil {
		return nil, err
	}

	// Clean up the version string
	return mcp.NewToolResultText(strings.TrimSpace(version)), nil
}

// GetBalanceSummary returns a summary of the balance for an address
type BalanceSummary struct {
	Address     string `json:"address"`
	TotalCoins  uint64 `json:"total_coins"`
	CoinCount   int    `json:"coin_count"`
	CoinObjects []any  `json:"coin_objects"`
}

// GetBalanceSummary returns a structured summary of the balance for an address
func (s *SuiService) GetBalanceSummary(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, ok := request.Params.Arguments["address"].(string)
	if !ok {
		return nil, errors.New("address must be a string")
	}
	output, err := s.client.GetBalance(address)
	if err != nil {
		return nil, err
	}

	// Try to parse the output as JSON
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("failed to parse balance output: %w", err)
	}

	// Extract coin objects
	coinObjects, ok := result["result"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format, missing result")
	}

	// Create a summary
	summary := &BalanceSummary{
		Address:     address,
		CoinObjects: []any{coinObjects},
	}

	// Additional processing could be done here

	return mcp.NewToolResultText(fmt.Sprintf("%v", summary)), nil
}

// GetObjectsSummary gets a summary of objects owned by an address
func (s *SuiService) GetObjectsSummary(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, ok := request.Params.Arguments["address"].(string)
	if !ok {
		return nil, errors.New("address must be a string")
	}
	output, err := s.client.GetObjects(address)
	if err != nil {
		return nil, err
	}

	// Try to parse the output as JSON
	var result interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("failed to parse objects output: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", result)), nil
}

// ProcessTransaction processes a transaction and returns readable information
func (s *SuiService) ProcessTransaction(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	txID, ok := request.Params.Arguments["txID"].(string)
	if !ok {
		return nil, errors.New("txID must be a string")
	}
	output, err := s.client.GetTransaction(txID)
	if err != nil {
		return nil, err
	}

	// Try to parse the output as JSON
	var result interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("failed to parse transaction output: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", result)), nil
}

// TransferTokens transfers tokens and returns the transaction result
func (s *SuiService) TransferTokens(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	recipient, ok := request.Params.Arguments["recipient"].(string)
	if !ok {
		return nil, errors.New("recipient must be a string")
	}
	amount, ok := request.Params.Arguments["amount"].(uint64)
	if !ok {
		return nil, errors.New("amount must be a uint64")
	}
	gasBudget, ok := request.Params.Arguments["gasBudget"].(string)
	if !ok {
		return nil, errors.New("gasBudget must be a string")
	}
	output, err := s.client.TransferSUI(recipient, amount, gasBudget)
	if err != nil {
		return nil, err
	}
	// Try to parse the output as JSON
	var result interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("failed to parse transaction output: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", result)), nil
}
