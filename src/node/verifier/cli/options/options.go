package options

import "github.com/c-bata/go-prompt"


var Initial = map[string]string{
	"eth":"ETH functions",
	"plasma":"ETH functions",
	"quit":"Quit from app",
}
var Eth = map[string]string{
	"transfer":"tr: Transfer -  eth tr 4 0x4ED6d26c6885247fA22746AB2c5328076597a5DF",
	"balance":"bal: Balance of eth account - eth bal 0x4ED6d26c6885247fA22746AB2c5328076597a5DF",
	"ownerBalance":"obal: Balance of user(config) eth account - eth obal",
}

var Plasma = map[string]string{
	"deposit":"dep: Deposit - plasma dep value",
	"balance":"bal: Plasma balance - plasma bal",
	"transfer":"tr: Transfer - plasma tr block txN out value address",
	"utxo":"utxo: List UTXOs from plasma - plasma utxo",
	"exit":"ex: Exit from plasma - plasma exit",
}


var InitialOptions = []prompt.Suggest{
	{Text: "eth", Description: Initial["eth"]},
	{Text: "plasma", Description: Initial["plasma"]},
	{Text: "quit", Description: Initial["quit"]},
}

var EthOptions = []prompt.Suggest{
	{Text: "transfer", Description: Eth["transfer"]},
	{Text: "balance", Description: Eth["balance"]},
	{Text: "ownerBalance", Description: Eth["ownerBalance"]},
}

var PlasmaOptions = []prompt.Suggest{
	{Text: "deposit", Description: Plasma["deposit"]},
	{Text: "balance", Description: Plasma["balance"]},
	{Text: "transfer", Description: Plasma["transfer"]},
	{Text: "utxo", Description: Plasma["utxo"]},
	{Text: "exit", Description: Plasma["exit"]},
	//{Text: "withdraw", Description: "ex: Withdraw eth from plasma"},
}
