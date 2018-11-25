package db

import "./leveldb"

func Tx(dbName string) (*leveldb.Table) {
	db, err := leveldb.Connect(dbName, 16,16,1024)
	txTable := leveldb.NewTable(db, "tx-")
	if err != nil {
		panic(err)
	}
	return txTable
}
