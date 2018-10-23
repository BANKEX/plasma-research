package database

import (
	"bytes"
	"errors"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

var utxoSS subspace.Subspace
var txSS subspace.Subspace

func Init() (fdb.Database, error) {
	fdb.MustAPIVersion(520)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"plasma"}, nil)
	if err != nil {
		return nil, err
	}

	utxoSS = schedulingDir.Sub("utxo")
	txSS = schedulingDir.Sub("tx")
	return db, nil
}

func ValidateInputs(tx Transaction) error {

}

func InsertTransaction(db fdb.Database, tx Transaction) error {
	ret, err := db.Transact(func(t fdb.Transaction) (interface{}, error) {
		/*
		READ /utxo/block:id:output (1-6x, read inputs for verification)
		DELETE /utxo/block:id:output (1-6x, spend inputs)
		READ /currentBlock
		INSERT /tx/currentBlock:hash (add transaction to the next block)
		 */

		currentBlock, err := t.Get(fdb.Key("currentBlock")).MustGet()
		_, err := t.Set(utxoSS.Pack(tuple.Tuple{currentBlock, hash}), tx)

		return
	})
	if err != nil {
		return err
	}
	if ret == nil {
		return errors.New("Could not write a transaction")
	}
	return nil
}