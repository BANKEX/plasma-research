package options

import "github.com/c-bata/go-prompt"

var InitialOptions = []prompt.Suggest{
	prompt.Suggest{Text: "plasma", Description: "Plasma functions"},
	prompt.Suggest{Text: "eth", Description: "ETH functions"},
	prompt.Suggest{Text: "main", Description: "Main functions"},
	prompt.Suggest{Text: "exit", Description: "Exit from app"},
}

var PlasmaOptions = []prompt.Suggest{
	prompt.Suggest{Text: "plasmaBalance", Description: "pb: Get balance of my account in Plasma"},
	prompt.Suggest{Text: "smartContractAddress", Description: "sca: Get Smart Contract address"},
	prompt.Suggest{Text: "smartContractBalance", Description: "scb: Get balance of Plasma smart contract"},
	prompt.Suggest{Text: "events", Description: "e: Get all events"},
}

var EthOptions = []prompt.Suggest{
	prompt.Suggest{Text: "smartContractAddress", Description: "sca: Get Smart Contract address"},
	prompt.Suggest{Text: "smartContractBalance", Description: "scb: Get balance of Plasma smart contract"},
	prompt.Suggest{Text: "transfer", Description: "tr: Transfer -  eth tr 0.4(eth) 0x4ED6..."},
	prompt.Suggest{Text: "accBalance", Description: "ab: Account Balance: eth ab 0x4ED6..."},
}

var MainOptions = []prompt.Suggest{
	prompt.Suggest{Text: "deposit", Description: "dep: Deposit"},
	prompt.Suggest{Text: "exit", Description: "ex: Exit from plasma"},
}
