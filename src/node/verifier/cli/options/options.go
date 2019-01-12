package options

import "github.com/c-bata/go-prompt"

var InitialOptions = []prompt.Suggest{
	{Text: "eth", Description: "ETH functions"},
	{Text: "plasma", Description: "Main functions"},
	{Text: "quit", Description: "Quit from app"},
}

var EthOptions = []prompt.Suggest{
	{Text: "transfer", Description: "tr: Transfer -  eth tr 0.4(eth) 0x4ED6..."},
	{Text: "balance", Description: "bal: Balance of eth account - eth bal 0x4ED6..."},
}

var PlasmaOptions = []prompt.Suggest{
	{Text: "deposit", Description: "dep: Deposit"},
	{Text: "transfer", Description: "tr: Transfer -  plasma tr 0.4(eth) 0x4ED6..."},
	{Text: "utxo", Description: "utxo: List UTXOs from plasma"},
	{Text: "exit", Description: "ex: Exit from plasma"},
	{Text: "withdraw", Description: "ex: Withdraw eth from plasma"},
}
