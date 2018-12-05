package blockchain

import (
	a "../../commons/alias"
	//"sync"
)

// TODO: better type naming
// Map string that looks like "BlockNumber:TransactionNumber:OutputNumber" to transaction bytes
type TxIndex map[string]a.TxHashBytes

type TransactionsPool map[a.TxHashBytes]Transaction

func (pool *TransactionsPool) IsEmpty() bool {
	return len(*pool) == 0
}

func map2slice(pool *TransactionsPool) []Transaction {
	list := make([]Transaction, 0, len(*pool))
	for _, t := range *pool {
		list = append(list, t)
	}
	return list
}

func (pool *TransactionsPool) GetTransactions() []Transaction {
	return map2slice(pool)
}

func (pool *TransactionsPool) Add(transaction Transaction) {
	var hash = transaction.GetHash()
	(*pool)[a.ToTxHashBytes(hash)] = transaction
}

func (pool *TransactionsPool) Remove(transaction Transaction) {
	var hash = transaction.GetHash()
	delete(*pool, a.ToTxHashBytes(hash))
}
