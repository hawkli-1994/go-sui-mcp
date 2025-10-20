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

func (s *SuiTools) GetSuiPath() mcp.Tool {
	return mcp.NewTool(
		"sui-path",
		mcp.WithDescription("Get the path of the local sui binary"),
	)
}

func (s *SuiTools) GetBalanceSummary() mcp.Tool {
	return mcp.NewTool(
		"sui-balance-summary",
		mcp.WithString("address",
			mcp.Description("Address to get the balance summary of, if not provided, the current address will be used"),
		),
		mcp.WithDescription("Get the balance summary of the Sui client"),
	)
}

func (s *SuiTools) GetObjectsSummary() mcp.Tool {
	return mcp.NewTool(
		"sui-objects-summary",
		mcp.WithString("address",
			mcp.Description("Address to get the objects summary of, if not provided, the current address will be used"),
		),
		mcp.WithDescription("Get the objects summary of the Sui client"),
	)
}

func (s *SuiTools) GetObject() mcp.Tool {
	return mcp.NewTool(
		"sui-object",
		mcp.WithString("objectID",
			mcp.Required(),
			mcp.Description("Object ID to get"),
		),
		mcp.WithDescription("Get the object of the Sui client"),
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

func (s *SuiTools) PaySUI() mcp.Tool {
	template := `交易状态：Success
	转账金额：1 SUI (1000000000 MIST)
	收款地址：0x398807039e4e99793c63a3a8b315c32c7878663e5f7ca0e9e19d3dddcbfb04f3
	交易 ID：2s7G1dNpVSfEM7uU1Zxy6aRqmYfgnDPv8cmFkg3aV89m
	Gas 费用：
	Storage Cost: 1976000 MIST
	Computation Cost: 1000000 MIST
	Storage Rebate: 978120 MIST
	Non-refundable Storage Fee: 9880 MIST
	转账已经完成，收款方已经收到了 1 SUI。新创建的 coin object ID 是：0x8bf92f132a6a9bd4ab32b083bcc0d54fc81bcd6fd07c0742c87e5c56e354cb04。`
	return mcp.NewTool(
		"sui-pay-sui",
		mcp.WithString("recipient",
			mcp.Required(),
			mcp.Description("Recipient address"),
		),
		mcp.WithNumber("amounts",
			mcp.Required(),
			mcp.Description("Amounts to transfer"),
		),
		mcp.WithString("gas-budget",
			mcp.Required(),
			mcp.Description("Gas budget"),
		),
		mcp.WithString("input-coins",
			mcp.Required(),
			mcp.Description("Input coins, need first to get the object of the SUI coin"),
		),

		mcp.WithDescription("Pay SUI, 首先, 第一步先检查整体余额是否足够, 然后调用objects summary 获取所有coin object, 然后调用object 找到余额足够的coin object, 然后调用pay sui 进行支付, 然后按照模板<"+template+">返回结果"),
	)
}

// ============ Address and Environment Management ============

func (s *SuiTools) GetActiveAddress() mcp.Tool {
	return mcp.NewTool(
		"sui-active-address",
		mcp.WithDescription("Get the current active address"),
	)
}

func (s *SuiTools) GetAddresses() mcp.Tool {
	return mcp.NewTool(
		"sui-addresses",
		mcp.WithDescription("List all addresses managed by the client"),
	)
}

func (s *SuiTools) GetActiveEnv() mcp.Tool {
	return mcp.NewTool(
		"sui-active-env",
		mcp.WithDescription("Get the current active environment (e.g., devnet, testnet, mainnet)"),
	)
}

func (s *SuiTools) GetEnvs() mcp.Tool {
	return mcp.NewTool(
		"sui-envs",
		mcp.WithDescription("List all Sui network environments configured"),
	)
}

func (s *SuiTools) GetChainIdentifier() mcp.Tool {
	return mcp.NewTool(
		"sui-chain-identifier",
		mcp.WithDescription("Query the chain identifier from the RPC endpoint"),
	)
}

// ============ Gas Management ============

func (s *SuiTools) GetGas() mcp.Tool {
	return mcp.NewTool(
		"sui-gas",
		mcp.WithString("address",
			mcp.Description("Address to get gas objects for, if not provided, the current address will be used"),
		),
		mcp.WithDescription("Get all gas objects owned by the address"),
	)
}

func (s *SuiTools) RequestFromFaucet() mcp.Tool {
	return mcp.NewTool(
		"sui-faucet",
		mcp.WithString("address",
			mcp.Description("Address to request gas coins for, if not provided, the current address will be used"),
		),
		mcp.WithDescription("Request gas coins from the faucet (works on devnet/testnet only)"),
	)
}

// ============ Transaction Operations ============

func (s *SuiTools) Transfer() mcp.Tool {
	return mcp.NewTool(
		"sui-transfer",
		mcp.WithString("to",
			mcp.Required(),
			mcp.Description("Recipient address"),
		),
		mcp.WithString("object-id",
			mcp.Required(),
			mcp.Description("Object ID to transfer"),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Transfer an object to another address"),
	)
}

func (s *SuiTools) TransferSUI() mcp.Tool {
	return mcp.NewTool(
		"sui-transfer-sui",
		mcp.WithString("to",
			mcp.Required(),
			mcp.Description("Recipient address"),
		),
		mcp.WithString("sui-coin-object-id",
			mcp.Required(),
			mcp.Description("SUI coin object ID to transfer from"),
		),
		mcp.WithNumber("amount",
			mcp.Description("Amount to transfer (in MIST). If not specified, transfers entire object"),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Transfer SUI to another address (simplified version)"),
	)
}

func (s *SuiTools) SplitCoin() mcp.Tool {
	return mcp.NewTool(
		"sui-split-coin",
		mcp.WithString("coin-id",
			mcp.Required(),
			mcp.Description("Coin object ID to split"),
		),
		mcp.WithArray("amounts",
			mcp.Required(),
			mcp.Description("Array of amounts to split into (in MIST)"),
			mcp.Items(map[string]interface{}{"type": "number"}),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Split a coin object into multiple coins"),
	)
}

func (s *SuiTools) MergeCoin() mcp.Tool {
	return mcp.NewTool(
		"sui-merge-coin",
		mcp.WithString("primary-coin",
			mcp.Required(),
			mcp.Description("Primary coin object ID (will receive merged balance)"),
		),
		mcp.WithString("coin-to-merge",
			mcp.Required(),
			mcp.Description("Coin object ID to merge into primary coin"),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Merge two coin objects into one"),
	)
}

func (s *SuiTools) Pay() mcp.Tool {
	return mcp.NewTool(
		"sui-pay",
		mcp.WithArray("input-coins",
			mcp.Required(),
			mcp.Description("Array of coin object IDs to use as input"),
			mcp.Items(map[string]interface{}{"type": "string"}),
		),
		mcp.WithArray("recipients",
			mcp.Required(),
			mcp.Description("Array of recipient addresses"),
			mcp.Items(map[string]interface{}{"type": "string"}),
		),
		mcp.WithArray("amounts",
			mcp.Required(),
			mcp.Description("Array of amounts to pay (in MIST)"),
			mcp.Items(map[string]interface{}{"type": "number"}),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Pay coins to multiple recipients with specified amounts"),
	)
}

func (s *SuiTools) PayAllSUI() mcp.Tool {
	return mcp.NewTool(
		"sui-pay-all-sui",
		mcp.WithArray("input-coins",
			mcp.Required(),
			mcp.Description("Array of SUI coin object IDs to use"),
			mcp.Items(map[string]interface{}{"type": "string"}),
		),
		mcp.WithString("recipient",
			mcp.Required(),
			mcp.Description("Recipient address"),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Pay all residual SUI to the recipient after deducting gas cost"),
	)
}

// ============ Contract Interaction ============

func (s *SuiTools) Call() mcp.Tool {
	return mcp.NewTool(
		"sui-call",
		mcp.WithString("package",
			mcp.Required(),
			mcp.Description("Package object ID"),
		),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
		mcp.WithString("function",
			mcp.Required(),
			mcp.Description("Function name to call"),
		),
		mcp.WithArray("type-args",
			mcp.Description("Type arguments for the function"),
			mcp.Items(map[string]interface{}{"type": "string"}),
		),
		mcp.WithArray("args",
			mcp.Description("Function arguments"),
			mcp.Items(map[string]interface{}{"type": "string"}),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for the transaction"),
		),
		mcp.WithDescription("Call a Move function on the Sui blockchain"),
	)
}

func (s *SuiTools) Publish() mcp.Tool {
	return mcp.NewTool(
		"sui-publish",
		mcp.WithString("package-path",
			mcp.Required(),
			mcp.Description("Path to the Move package directory"),
		),
		mcp.WithString("gas-budget",
			mcp.Description("Gas budget for publishing"),
		),
		mcp.WithBoolean("skip-dependency-verification",
			mcp.Description("Skip dependency verification"),
		),
		mcp.WithDescription("Publish Move modules to the Sui blockchain"),
	)
}

func (s *SuiTools) GetDynamicField() mcp.Tool {
	return mcp.NewTool(
		"sui-dynamic-field",
		mcp.WithString("parent-object-id",
			mcp.Required(),
			mcp.Description("Parent object ID that contains the dynamic field"),
		),
		mcp.WithString("name",
			mcp.Description("Name of the dynamic field to query"),
		),
		mcp.WithDescription("Query a dynamic field by its parent object address"),
	)
}

// ============ Move Development ============

func (s *SuiTools) MoveBuild() mcp.Tool {
	return mcp.NewTool(
		"sui-move-build",
		mcp.WithString("package-path",
			mcp.Description("Path to the Move package (defaults to current directory)"),
		),
		mcp.WithDescription("Build a Move package"),
	)
}

func (s *SuiTools) MoveTest() mcp.Tool {
	return mcp.NewTool(
		"sui-move-test",
		mcp.WithString("package-path",
			mcp.Description("Path to the Move package (defaults to current directory)"),
		),
		mcp.WithString("filter",
			mcp.Description("Filter tests by name pattern"),
		),
		mcp.WithDescription("Run Move unit tests in the package"),
	)
}

func (s *SuiTools) MoveNew() mcp.Tool {
	return mcp.NewTool(
		"sui-move-new",
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the new Move package"),
		),
		mcp.WithString("path",
			mcp.Description("Path where to create the package (defaults to ./name)"),
		),
		mcp.WithDescription("Create a new Move package"),
	)
}

// ============ Keytool Management ============

func (s *SuiTools) KeytoolList() mcp.Tool {
	return mcp.NewTool(
		"sui-keytool-list",
		mcp.WithDescription("List all keys in the Sui keystore"),
	)
}

func (s *SuiTools) KeytoolGenerate() mcp.Tool {
	return mcp.NewTool(
		"sui-keytool-generate",
		mcp.WithString("key-scheme",
			mcp.Required(),
			mcp.Description("Key scheme: ed25519, secp256k1, or secp256r1"),
		),
		mcp.WithString("derivation-path",
			mcp.Description("Derivation path (e.g., m/44'/784'/0'/0'/0')"),
		),
		mcp.WithString("word-length",
			mcp.Description("Word length for mnemonic: word12, word15, word18, word21, or word24"),
		),
		mcp.WithDescription("Generate a new keypair and add it to the keystore"),
	)
}

func (s *SuiTools) KeytoolExport() mcp.Tool {
	return mcp.NewTool(
		"sui-keytool-export",
		mcp.WithString("address",
			mcp.Required(),
			mcp.Description("Address or alias of the key to export"),
		),
		mcp.WithDescription("Export the private key for a given address (Bech32 encoded)"),
	)
}
