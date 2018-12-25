package options

import "github.com/c-bata/go-prompt"

var InitialOptions = []prompt.Suggest{
	{Text: "plasma", Description: "Plasma functions"},
	{Text: "eth", Description: "ETH functions"},
	{Text: "main", Description: "Main functions"},
	{Text: "exit", Description: "Exit from app"},
}

var PlasmaOptions = []prompt.Suggest{
	{Text: "plasmaBalance", Description: "pb: Get balance of my account in Plasma"},
	{Text: "smartContractAddress", Description: "sca: Get Smart Contract address"},
	{Text: "smartContractBalance", Description: "scb: Get balance of Plasma smart contract"},
	{Text: "events", Description: "e: Get all events"},
}

var EthOptions = []prompt.Suggest{
	{Text: "smartContractAddress", Description: "sca: Get Smart Contract address"},
	{Text: "smartContractBalance", Description: "scb: Get balance of Plasma smart contract"},
	{Text: "transfer", Description: "tr: Transfer -  eth tr 0.4(eth) 0x4ED6..."},
	{Text: "accBalance", Description: "ab: Account Balance: eth ab 0x4ED6..."},
}

var MainOptions = []prompt.Suggest{
	{Text: "deposit", Description: "dep: Deposit"},
	{Text: "exit", Description: "ex: Exit from plasma"},
}
