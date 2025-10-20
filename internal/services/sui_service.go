package services

import (
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

func (s *SuiService) GetSuiPath(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := s.client.GetSuiPath()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(path), nil
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
	address, _ := request.Params.Arguments["address"].(string)
	output, err := s.client.GetBalance(address)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(output), nil
}

// GetObjectsSummary gets a summary of objects owned by an address
func (s *SuiService) GetObjectsSummary(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, _ := request.Params.Arguments["address"].(string)

	output, err := s.client.GetObjects(address)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(output), nil
}

// GetObject processes a transaction and returns readable information
func (s *SuiService) GetObject(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	objectID, ok := request.Params.Arguments["objectID"].(string)
	if !ok {
		return nil, errors.New("objectID must be a string")
	}
	output, err := s.client.GetObject(objectID)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(output), nil
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

	return mcp.NewToolResultText(output), nil
}

// PaySUI transfers tokens and returns the transaction result
func (s *SuiService) PaySUI(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	recipient, ok := request.Params.Arguments["recipient"].(string)
	if !ok {
		return nil, errors.New("recipient must be a string")
	}
	amountFloat, ok := request.Params.Arguments["amounts"].(float64)
	if !ok {
		return nil, errors.New("amounts must be a number")
	}
	amounts := uint64(amountFloat)

	inputCoins, ok := request.Params.Arguments["input-coins"].(string)
	if !ok {
		return nil, errors.New("inputCoins must be a string")
	}

	gasBudget, ok := request.Params.Arguments["gas-budget"].(string)
	if !ok {
		return nil, errors.New("gasBudget must be a string")
	}
	output, err := s.client.PaySUI(recipient, inputCoins, amounts, gasBudget)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(output), nil
}

// ============ Address and Environment Management ============

// GetActiveAddress returns the current active address
func (s *SuiService) GetActiveAddress(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.GetActiveAddress()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strings.TrimSpace(output)), nil
}

// GetAddresses returns all addresses managed by the client
func (s *SuiService) GetAddresses(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.GetAddresses()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// GetActiveEnv returns the current active environment
func (s *SuiService) GetActiveEnv(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.GetActiveEnv()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strings.TrimSpace(output)), nil
}

// GetEnvs returns all Sui environments
func (s *SuiService) GetEnvs(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.GetEnvs()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// GetChainIdentifier queries the chain identifier from the RPC endpoint
func (s *SuiService) GetChainIdentifier(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.GetChainIdentifier()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strings.TrimSpace(output)), nil
}

// ============ Gas Management ============

// GetGas obtains all gas objects owned by the address
func (s *SuiService) GetGas(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, _ := request.Params.Arguments["address"].(string)
	output, err := s.client.GetGas(address)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// RequestFromFaucet requests gas coins from faucet
func (s *SuiService) RequestFromFaucet(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, _ := request.Params.Arguments["address"].(string)
	output, err := s.client.RequestFromFaucet(address)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// ============ Transaction Operations ============

// Transfer transfers an object to another address
func (s *SuiService) Transfer(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	to, ok := request.Params.Arguments["to"].(string)
	if !ok {
		return nil, errors.New("to must be a string")
	}
	objectID, ok := request.Params.Arguments["object-id"].(string)
	if !ok {
		return nil, errors.New("object-id must be a string")
	}
	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.Transfer(to, objectID, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// TransferSUI transfers SUI to another address
func (s *SuiService) TransferSUI(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	to, ok := request.Params.Arguments["to"].(string)
	if !ok {
		return nil, errors.New("to must be a string")
	}
	suiCoinObjectID, ok := request.Params.Arguments["sui-coin-object-id"].(string)
	if !ok {
		return nil, errors.New("sui-coin-object-id must be a string")
	}

	var amount uint64
	if amountFloat, ok := request.Params.Arguments["amount"].(float64); ok {
		amount = uint64(amountFloat)
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.TransferSUI(to, suiCoinObjectID, amount, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// SplitCoin splits a coin object into multiple coins
func (s *SuiService) SplitCoin(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	coinID, ok := request.Params.Arguments["coin-id"].(string)
	if !ok {
		return nil, errors.New("coin-id must be a string")
	}

	amountsInterface, ok := request.Params.Arguments["amounts"].([]interface{})
	if !ok {
		return nil, errors.New("amounts must be an array")
	}

	amounts := make([]uint64, len(amountsInterface))
	for i, v := range amountsInterface {
		if num, ok := v.(float64); ok {
			amounts[i] = uint64(num)
		} else {
			return nil, errors.New("amounts must contain numbers")
		}
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.SplitCoin(coinID, amounts, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// MergeCoin merges two coin objects into one
func (s *SuiService) MergeCoin(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	primaryCoin, ok := request.Params.Arguments["primary-coin"].(string)
	if !ok {
		return nil, errors.New("primary-coin must be a string")
	}
	coinToMerge, ok := request.Params.Arguments["coin-to-merge"].(string)
	if !ok {
		return nil, errors.New("coin-to-merge must be a string")
	}
	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.MergeCoin(primaryCoin, coinToMerge, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// Pay pays coins to recipients following specified amounts
func (s *SuiService) Pay(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	inputCoinsInterface, ok := request.Params.Arguments["input-coins"].([]interface{})
	if !ok {
		return nil, errors.New("input-coins must be an array")
	}
	inputCoins := make([]string, len(inputCoinsInterface))
	for i, v := range inputCoinsInterface {
		if str, ok := v.(string); ok {
			inputCoins[i] = str
		} else {
			return nil, errors.New("input-coins must contain strings")
		}
	}

	recipientsInterface, ok := request.Params.Arguments["recipients"].([]interface{})
	if !ok {
		return nil, errors.New("recipients must be an array")
	}
	recipients := make([]string, len(recipientsInterface))
	for i, v := range recipientsInterface {
		if str, ok := v.(string); ok {
			recipients[i] = str
		} else {
			return nil, errors.New("recipients must contain strings")
		}
	}

	amountsInterface, ok := request.Params.Arguments["amounts"].([]interface{})
	if !ok {
		return nil, errors.New("amounts must be an array")
	}
	amounts := make([]uint64, len(amountsInterface))
	for i, v := range amountsInterface {
		if num, ok := v.(float64); ok {
			amounts[i] = uint64(num)
		} else {
			return nil, errors.New("amounts must contain numbers")
		}
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.Pay(inputCoins, recipients, amounts, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// PayAllSUI pays all residual SUI coins to the recipient
func (s *SuiService) PayAllSUI(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	inputCoinsInterface, ok := request.Params.Arguments["input-coins"].([]interface{})
	if !ok {
		return nil, errors.New("input-coins must be an array")
	}
	inputCoins := make([]string, len(inputCoinsInterface))
	for i, v := range inputCoinsInterface {
		if str, ok := v.(string); ok {
			inputCoins[i] = str
		} else {
			return nil, errors.New("input-coins must contain strings")
		}
	}

	recipient, ok := request.Params.Arguments["recipient"].(string)
	if !ok {
		return nil, errors.New("recipient must be a string")
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.PayAllSUI(inputCoins, recipient, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// ============ Contract Interaction ============

// Call calls a Move function
func (s *SuiService) Call(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	packageID, ok := request.Params.Arguments["package"].(string)
	if !ok {
		return nil, errors.New("package must be a string")
	}
	module, ok := request.Params.Arguments["module"].(string)
	if !ok {
		return nil, errors.New("module must be a string")
	}
	function, ok := request.Params.Arguments["function"].(string)
	if !ok {
		return nil, errors.New("function must be a string")
	}

	var typeArgs []string
	if typeArgsInterface, ok := request.Params.Arguments["type-args"].([]interface{}); ok {
		typeArgs = make([]string, len(typeArgsInterface))
		for i, v := range typeArgsInterface {
			if str, ok := v.(string); ok {
				typeArgs[i] = str
			}
		}
	}

	var args []string
	if argsInterface, ok := request.Params.Arguments["args"].([]interface{}); ok {
		args = make([]string, len(argsInterface))
		for i, v := range argsInterface {
			if str, ok := v.(string); ok {
				args[i] = str
			}
		}
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)

	output, err := s.client.Call(packageID, module, function, typeArgs, args, gasBudget)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// Publish publishes Move modules
func (s *SuiService) Publish(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	packagePath, ok := request.Params.Arguments["package-path"].(string)
	if !ok {
		return nil, errors.New("package-path must be a string")
	}

	gasBudget, _ := request.Params.Arguments["gas-budget"].(string)
	skipDependencyVerification, _ := request.Params.Arguments["skip-dependency-verification"].(bool)

	output, err := s.client.Publish(packagePath, gasBudget, skipDependencyVerification)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// GetDynamicField queries a dynamic field by its address
func (s *SuiService) GetDynamicField(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	parentObjectID, ok := request.Params.Arguments["parent-object-id"].(string)
	if !ok {
		return nil, errors.New("parent-object-id must be a string")
	}

	name, _ := request.Params.Arguments["name"].(string)

	output, err := s.client.GetDynamicField(parentObjectID, name)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// ============ Move Development ============

// MoveBuild builds a Move package
func (s *SuiService) MoveBuild(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	packagePath, _ := request.Params.Arguments["package-path"].(string)

	output, err := s.client.MoveBuild(packagePath)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// MoveTest runs Move unit tests
func (s *SuiService) MoveTest(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	packagePath, _ := request.Params.Arguments["package-path"].(string)
	filter, _ := request.Params.Arguments["filter"].(string)

	output, err := s.client.MoveTest(packagePath, filter)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// MoveNew creates a new Move package
func (s *SuiService) MoveNew(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	path, _ := request.Params.Arguments["path"].(string)

	output, err := s.client.MoveNew(name, path)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// ============ Keytool Management ============

// KeytoolList lists all keys in the keystore
func (s *SuiService) KeytoolList(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	output, err := s.client.KeytoolList()
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// KeytoolGenerate generates a new keypair
func (s *SuiService) KeytoolGenerate(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyScheme, ok := request.Params.Arguments["key-scheme"].(string)
	if !ok {
		return nil, errors.New("key-scheme must be a string")
	}

	derivationPath, _ := request.Params.Arguments["derivation-path"].(string)
	wordLength, _ := request.Params.Arguments["word-length"].(string)

	output, err := s.client.KeytoolGenerate(keyScheme, derivationPath, wordLength)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}

// KeytoolExport exports the private key for a given address
func (s *SuiService) KeytoolExport(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	address, ok := request.Params.Arguments["address"].(string)
	if !ok {
		return nil, errors.New("address must be a string")
	}

	output, err := s.client.KeytoolExport(address)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(output), nil
}
