package sui

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

// Client provides methods to interact with the Sui client CLI
type Client struct {
	executablePath string
}

// NewClient creates a new Sui client instance
func NewClient() *Client {
	// Use the configured executable path or default to "sui"
	execPath := viper.GetString("sui.executable_path")
	if execPath == "" {
		execPath = "sui"
	}

	return &Client{
		executablePath: execPath,
	}
}

// ExecuteCommand runs a Sui command and returns the output
func (c *Client) ExecuteCommand(args ...string) (string, error) {
	cmd := exec.Command(c.executablePath, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing sui command: %v\nStderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

// GetVersion returns the Sui client version
func (c *Client) GetVersion() (string, error) {
	output, err := c.ExecuteCommand("--version")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

func (c *Client) GetSuiPath() (string, error) {

	output, err := exec.Command("which", "sui").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// GetBalance gets the balance for a specific address
func (c *Client) GetBalance(address string) (string, error) {
	args := []string{"client", "balance"}
	if address != "" {
		args = append(args, address)
	}
	return c.ExecuteCommand(args...)
}

// GetObjects gets objects owned by an address
func (c *Client) GetObjects(address string) (string, error) {
	args := []string{"client", "objects"}
	if address != "" {
		args = append(args, address)
	}
	args = append(args, "--json")
	return c.ExecuteCommand(args...)
}

func (c *Client) GetObject(objectID string) (string, error) {
	args := []string{"client", "object", objectID, "--json"}
	return c.ExecuteCommand(args...)
}

// GetActiveValidators gets the list of active validators
func (c *Client) GetActiveValidators() (string, error) {
	args := []string{"client", "active-validators"}
	return c.ExecuteCommand(args...)
}

// GetNetwork returns the current network info
func (c *Client) GetNetwork() (string, error) {
	args := []string{"client", "envs"}
	return c.ExecuteCommand(args...)
}

// GetTransaction retrieves information about a specific transaction
func (c *Client) GetTransaction(txID string) (string, error) {
	args := []string{"client", "tx-block", txID}
	return c.ExecuteCommand(args...)
}

// PaySUI transfers SUI tokens to recipients (supports multiple recipients)
func (c *Client) PaySUI(recipients []string, inputCoins []string, amounts []uint64, gasBudget string) (string, error) {
	args := []string{"client", "pay-sui"}

	// Add multiple --input-coins flags
	for _, coin := range inputCoins {
		args = append(args, "--input-coins", coin)
	}

	// Add multiple --recipients flags
	for _, recipient := range recipients {
		args = append(args, "--recipients", recipient)
	}

	// Add multiple --amounts flags
	for _, amount := range amounts {
		args = append(args, "--amounts", fmt.Sprintf("%d", amount))
	}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// ============ Address and Environment Management ============

// GetActiveAddress returns the current active address
func (c *Client) GetActiveAddress() (string, error) {
	args := []string{"client", "active-address"}
	return c.ExecuteCommand(args...)
}

// GetAddresses returns all addresses managed by the client
func (c *Client) GetAddresses() (string, error) {
	args := []string{"client", "addresses"}
	return c.ExecuteCommand(args...)
}

// GetActiveEnv returns the current active environment
func (c *Client) GetActiveEnv() (string, error) {
	args := []string{"client", "active-env"}
	return c.ExecuteCommand(args...)
}

// GetEnvs returns all Sui environments
func (c *Client) GetEnvs() (string, error) {
	args := []string{"client", "envs"}
	return c.ExecuteCommand(args...)
}

// GetChainIdentifier queries the chain identifier from the RPC endpoint
func (c *Client) GetChainIdentifier() (string, error) {
	args := []string{"client", "chain-identifier"}
	return c.ExecuteCommand(args...)
}

// ============ Gas Management ============

// GetGas obtains all gas objects owned by the address
func (c *Client) GetGas(address string) (string, error) {
	args := []string{"client", "gas"}
	if address != "" {
		args = append(args, address)
	}
	return c.ExecuteCommand(args...)
}

// RequestFromFaucet requests gas coins from faucet
func (c *Client) RequestFromFaucet(address string) (string, error) {
	args := []string{"client", "faucet"}
	if address != "" {
		args = append(args, "--address", address)
	}
	return c.ExecuteCommand(args...)
}

// ============ Transaction Operations ============

// Transfer transfers an object to another address
func (c *Client) Transfer(to string, objectID string, gasBudget string) (string, error) {
	args := []string{"client", "transfer",
		"--to", to,
		"--object-id", objectID}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// TransferSUI transfers SUI to another address
func (c *Client) TransferSUI(to string, suiCoinObjectID string, amount uint64, gasBudget string) (string, error) {
	args := []string{"client", "transfer-sui",
		"--to", to,
		"--sui-coin-object-id", suiCoinObjectID}

	if amount > 0 {
		args = append(args, "--amount", fmt.Sprintf("%d", amount))
	}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// SplitCoin splits a coin object into multiple coins
func (c *Client) SplitCoin(coinID string, amounts []uint64, gasBudget string) (string, error) {
	args := []string{"client", "split-coin",
		"--coin-id", coinID}

	// Convert amounts to comma-separated string
	amountStrs := make([]string, len(amounts))
	for i, amt := range amounts {
		amountStrs[i] = fmt.Sprintf("%d", amt)
	}
	args = append(args, "--amounts", strings.Join(amountStrs, ","))

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// MergeCoin merges two coin objects into one
func (c *Client) MergeCoin(primaryCoin string, coinToMerge string, gasBudget string) (string, error) {
	args := []string{"client", "merge-coin",
		"--primary-coin", primaryCoin,
		"--coin-to-merge", coinToMerge}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// Pay pays coins to recipients following specified amounts
func (c *Client) Pay(inputCoins []string, recipients []string, amounts []uint64, gasBudget string) (string, error) {
	args := []string{"client", "pay"}

	// Add multiple --input-coins flags
	for _, coin := range inputCoins {
		args = append(args, "--input-coins", coin)
	}

	// Add multiple --recipients flags
	for _, recipient := range recipients {
		args = append(args, "--recipients", recipient)
	}

	// Add multiple --amounts flags
	for _, amount := range amounts {
		args = append(args, "--amounts", fmt.Sprintf("%d", amount))
	}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// PayAllSUI pays all residual SUI coins to the recipient
func (c *Client) PayAllSUI(inputCoins []string, recipient string, gasBudget string) (string, error) {
	args := []string{"client", "pay-all-sui"}

	// Add multiple --input-coins flags
	for _, coin := range inputCoins {
		args = append(args, "--input-coins", coin)
	}

	args = append(args, "--recipient", recipient)

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(args...)
}

// ============ Contract Interaction ============

// Call calls a Move function
func (c *Client) Call(packageID string, module string, function string, typeArgs []string, args []string, gasBudget string) (string, error) {
	cmdArgs := []string{"client", "call",
		"--package", packageID,
		"--module", module,
		"--function", function}

	// Add multiple --type-args flags
	for _, typeArg := range typeArgs {
		cmdArgs = append(cmdArgs, "--type-args", typeArg)
	}

	// Add multiple --args flags
	for _, arg := range args {
		cmdArgs = append(cmdArgs, "--args", arg)
	}

	if gasBudget != "" {
		cmdArgs = append(cmdArgs, "--gas-budget", gasBudget)
	}

	return c.ExecuteCommand(cmdArgs...)
}

// Publish publishes Move modules
func (c *Client) Publish(packagePath string, gasBudget string, skipDependencyVerification bool) (string, error) {
	args := []string{"client", "publish", packagePath}

	if gasBudget != "" {
		args = append(args, "--gas-budget", gasBudget)
	}

	if skipDependencyVerification {
		args = append(args, "--skip-dependency-verification")
	}

	return c.ExecuteCommand(args...)
}

// GetDynamicField queries a dynamic field by its address
func (c *Client) GetDynamicField(parentObjectID string, name string) (string, error) {
	args := []string{"client", "dynamic-field", parentObjectID}
	if name != "" {
		args = append(args, "--name", name)
	}
	return c.ExecuteCommand(args...)
}

// ============ Move Development ============

// MoveBuild builds a Move package
func (c *Client) MoveBuild(packagePath string) (string, error) {
	args := []string{"move", "build"}
	if packagePath != "" {
		args = append(args, "--path", packagePath)
	}
	return c.ExecuteCommand(args...)
}

// MoveTest runs Move unit tests
func (c *Client) MoveTest(packagePath string, filter string) (string, error) {
	args := []string{"move", "test"}
	if packagePath != "" {
		args = append(args, "--path", packagePath)
	}
	if filter != "" {
		args = append(args, "--filter", filter)
	}
	return c.ExecuteCommand(args...)
}

// MoveNew creates a new Move package
func (c *Client) MoveNew(name string, path string) (string, error) {
	args := []string{"move", "new", name}
	if path != "" {
		args = append(args, path)
	}
	return c.ExecuteCommand(args...)
}

// ============ Keytool Management ============

// KeytoolList lists all keys in the keystore
func (c *Client) KeytoolList() (string, error) {
	args := []string{"keytool", "list"}
	return c.ExecuteCommand(args...)
}

// KeytoolGenerate generates a new keypair
func (c *Client) KeytoolGenerate(keyScheme string, derivationPath string, wordLength string) (string, error) {
	args := []string{"keytool", "generate", keyScheme}
	if derivationPath != "" {
		args = append(args, "--derivation-path", derivationPath)
	}
	if wordLength != "" {
		args = append(args, "--word-length", wordLength)
	}
	return c.ExecuteCommand(args...)
}

// KeytoolExport exports the private key for a given address
func (c *Client) KeytoolExport(address string) (string, error) {
	args := []string{"keytool", "export", address}
	return c.ExecuteCommand(args...)
}

// ============ Helper Functions ============

// uint64SliceToStrings converts a slice of uint64 to a slice of strings
func uint64SliceToStrings(nums []uint64) []string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = fmt.Sprintf("%d", num)
	}
	return strs
}
