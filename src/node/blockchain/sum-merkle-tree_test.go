package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSMT(t *testing.T) {
	data := []byte(`
	[
		{
			"inputs": [
				{
					"blockNumber":1,
					"txNumber":1,
					"outputNumber":1,
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 1,
						"end": 2
					}
				}

			],
			"outputs": [
				{
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 1,
						"end": 2
					}
				}
			],
			"metadata": {
				"maxBlockNumber":100
			}
		},
		{
			"inputs": [
				{
					"blockNumber":1,
					"txNumber":1,
					"outputNumber":1,
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 4,
						"end": 5
					}
				}

			],
			"outputs": [
				{
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 4,
						"end": 5
					}
				}
			],
			"metadata": {
				"maxBlockNumber":100
			}
		}
	]
	`)

	q := make([]Transaction, 0)
	json.Unmarshal(data, &q)
	var sumTree *SumMerkleTree

	txs, err := PrepareLeaves(q)
	if err != nil {
		assert.Error(t, err)
	}

	sumTree = NewSumMerkleTree(txs)

	root := sumTree.GetRoot()
	fmt.Printf("rootHash = %x\n", root.Hash)
	fmt.Printf("rootLength = %x\n", root.Length)

	leaf := sumTree.Leafs[1];
	fmt.Printf("Begin = %d\n", leaf.Begin)
	fmt.Printf("End = %d\n", leaf.End)
	proof := sumTree.GetProof(1)

	fmt.Printf("Item=%x\n", proof.Item)
	for _, item := range proof.Data {
		fmt.Printf("%x%x\n", item.Length, item.Hash)
	}
}
