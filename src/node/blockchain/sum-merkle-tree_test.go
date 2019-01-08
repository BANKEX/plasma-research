package blockchain

import (
	"encoding/json"
	"fmt"
	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
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

	leaf := sumTree.Leafs[1]
	fmt.Printf("Begin = %d\n", leaf.Begin)
	fmt.Printf("End = %d\n", leaf.End)
	proof := sumTree.GetProof(1)

	fmt.Printf("Item=%x\n", proof.Item)
	for _, item := range proof.Data {
		fmt.Printf("%x%x\n", item.Length, item.Hash)
	}

	fmt.Printf("RlpEncoded=%x\n", sumTree.GetRlpEncodedProof(1))
}

func TestFillGapsOneSlice(t *testing.T) {

	oneSlice := func(b uint32, e uint32) string {
		slices := []Slice{{Begin: b, End: e}}
		return fmt.Sprint(FillGaps(slices))
	}

	// Slice at the beginning
	assert.Equal(t, "[{0 1} {1 16777215}]", oneSlice(0, 1))
	// Slice in the middle
	assert.Equal(t, "[{0 1000} {1000 2000} {2000 16777215}]", oneSlice(1000, 2000))
	// Slice at the end
	assert.Equal(t, "[{0 16777200} {16777200 16777215}]", oneSlice(16777200, 16777215))

	///
	twoSlices := func(b1 uint32, e1 uint32, b2 uint32, e2 uint32) string {
		slices := []Slice{
			{Begin: b1, End: e1},
			{Begin: b2, End: e2},
		}
		return fmt.Sprint(FillGaps(slices))
	}

	// Slices at the beginning and at the end
	assert.Equal(t, "[{0 10} {10 16777000} {16777000 16777215}]", twoSlices(0, 10, 16777000, 16777215))

	// Two Slices in the middle
	assert.Equal(t, "[{0 100} {100 200} {200 500} {500 600} {600 16777215}]", twoSlices(100, 200, 500, 600))

	// Fill gaps between three slices
	threeSlices := []Slice{
		{0, 10},
		{200, 500},
		{3000, 16777215},
	}
	assert.Equal(t, "[{0 10} {10 200} {200 500} {500 3000} {3000 16777215}]",
		fmt.Sprint(FillGaps(threeSlices)))

	// Return just one slice is source collection is empty
	assert.Equal(t, "[{0 16777215}]", fmt.Sprint(FillGaps([]Slice{})))
}
