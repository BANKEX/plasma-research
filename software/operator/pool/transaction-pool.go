package pool

import (
	a "../../commons/alias"
	b "../../commons/blockchain"
	//"sync"
)

type TransactionsPool map[a.TxHashBytes]b.Transaction

//var instance *TransactionsPool
//var once sync.Once
//
//func GetTransactionsPool() *TransactionsPool {
//	once.Do(func() {
//		instance = &TransactionsPool{}
//	})
//	return instance
//}

//func GetTransactionsPoolCopy() *TransactionsPool{
//	var poolCopy TransactionsPool
//	//pool := GetTransactionsPool()
//	for k,v := range *pool{
//		poolCopy[k] = v
//	}
//	return &poolCopy
//}

func (pool *TransactionsPool) IsEmpty() bool {
	return len(*pool) == 0
}

func map2slice(pool *TransactionsPool) []b.Transaction {
	list := make([]b.Transaction, 0, len(*pool))
	for _, t := range *pool {
		list = append(list, t)
	}
	return list
}

func (pool *TransactionsPool) GetTransactions() []b.Transaction {
	return map2slice(pool)
}

func (pool *TransactionsPool) Add(transaction b.Transaction) {
	var hash = transaction.GetHash()
	(*pool)[a.ToTxHashBytes(hash)] = transaction
}

func (pool *TransactionsPool) Remove(transaction b.Transaction) {
	var hash = transaction.GetHash()
	delete(*pool, a.ToTxHashBytes(hash))
}
