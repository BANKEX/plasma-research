package blockchain

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/stretchr/testify/assert"
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
	fmt.Println()
}

func TestFirstElement(t *testing.T) {
	slice := []Slice{Slice{Begin: 1, End: 2}}
	res := FillGapsWithSlices(slice)

	// [{0 1} {1 2} {2 16777215]}]
	fmt.Println("TestFirstElement:")
	fmt.Println(res)

	if res[0].Begin != 0 || res[0].End != 1 {
		t.Error(res[0])
		fmt.Println("True: [{0 1} {1 2} {2 16777215]}]")
	}

	fmt.Println()
}

func TestLastElement(t *testing.T) {
	slice := []Slice{Slice{Begin: 2, End: 4}}
	res := FillGapsWithSlices(slice)

	// [{0 2} {2 4} {4 16777215}]
	fmt.Println("TestLastElement:")
	fmt.Println(res)

	if res[len(res)-1].Begin != 4 || res[len(res)-1].End != plasmaLength {
		fmt.Println(res[len(res)-1])
		t.Error(res[len(res)-1])
		fmt.Println("True: [{0 2} {2 4} {4 16777215}]")
	}
	fmt.Println()
}

func TestMiddleElement(t *testing.T) {
	slice := []Slice{Slice{Begin: 0, End: 1}, Slice{Begin: 2, End: 4}}
	res := FillGapsWithSlices(slice)

	// [{0 1} {1 2} {2 4} {4 16777215}]
	fmt.Println("TestMiddleElement:")
	fmt.Println(res)

	if res[1].Begin != 1 || res[1].End != 6 {
		t.Error(res[1])
		fmt.Println("True: [{0 1} {1 2} {2 4} {4 16777215}]")
	}
	fmt.Println()
}

func TestFirstAndLastElement(t *testing.T) {
	slice := []Slice{Slice{Begin: 8, End: 10}, Slice{Begin: 15, End: 20}}
	res := FillGapsWithSlices(slice)

	// [{0 8} {8 10} {10 15} {15 20} {20 16777215}]

	fmt.Println("TestFirstAndLastElement:")
	fmt.Println(res)

	if res[0].Begin != 0 || res[0].End != 8 {
		t.Error(res[0])
		fmt.Println("True: [{0 8} {8 10} {10 15} {15 20} {20 16777215}]")
	}

	if res[len(res)-1].Begin != 20 || res[len(res)-1].End != plasmaLength {
		fmt.Println(res[len(res)-1])
		fmt.Println("True: [{0 8} {8 10} {10 15} {15 20} {20 16777215}]")
	}
	fmt.Println()
}

func TestAllElements(t *testing.T) {
	slice := []Slice{Slice{Begin: 2, End: 3}, Slice{Begin: 100, End: 300}}
	res := FillGapsWithSlices(slice)

	// [{0 2} {2 3} {3 100} {100 300} {300 16777215}]

	fmt.Println("TestAllElements:")
	fmt.Println(res)

	// 1
	if res[0].Begin != 0 || res[0].End != 2 {
		t.Error(res[0])
		fmt.Println("True: [{0 2} {2 3} {3 100} {100 300} {300 16777215}]")
	}

	// 2
	if res[1].Begin != 2 || res[1].End != 3 {
		t.Error(res[1])
		fmt.Println("True: [{0 2} {2 3} {3 100} {100 300} {300 16777215}]")
	}

	// 3
	if res[2].Begin != 3 || res[2].End != 100 {
		t.Error(res[2])
		fmt.Println("True: [{0 2} {2 3} {3 100} {100 300} {300 16777215}]")
	}

	// 4
	if res[3].Begin != 100 || res[3].End != 300 {
		t.Error(res[3])
		fmt.Println("True: [{0 2} {2 3} {3 100} {100 300} {300 16777215}]")
	}

	// 5
	if res[len(res)-1].Begin != 300 || res[len(res)-1].End != plasmaLength {
		t.Error(res[len(res)-1])
		fmt.Println("True: [{0 2} {2 3} {3 100} {100 300} {300 16777215}]")
	}
	fmt.Println()
}
