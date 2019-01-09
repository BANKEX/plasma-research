package blockchain

import (
	"encoding/json"
	"fmt"
	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestSMT(t *testing.T) {

	type Data struct {
		Info []struct {
			Inputs []struct {
				BlockNumber  int   `json:"blockNumber"`
				TxNumber     int   `json:"txNumber"`
				OutputNumber int   `json:"outputNumber"`
				Owner        []int `json:"owner"`
				Slice        struct {
					Begin int `json:"begin"`
					End   int `json:"end"`
				}
			}
			Outputs []struct {
				Owner []int `json:"owner"`
				Slice struct {
					Begin int `json:"begin"`
					End   int `json:"end"`
				}
			}
			Metadata struct {
				MaxBlockNumber int `json:"maxBlockNumber"`
			}
		}
	}

	jsonFile, err := os.Open("conf.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	var d Data
	json.Unmarshal(byteValue, &d)

	data, err := json.Marshal(d.Info)
	if err != nil {
		fmt.Println(err)
	}

	type Result struct {
		RootHash        string   `json:"rootHash"`
		RootLength      string   `json:"rootLength"`
		Begin           string   `json:"begin"`
		End             string   `json:"end"`
		Item            string   `json:"item"`
		ItemsLenAndHash []string `json:"itemsLenAndHash"`
		RlpEncoded      string   `json:"rlpEncoded"`
	}

	var result Result

	q := make([]Transaction, 0)
	json.Unmarshal(data, &q)

	var sumTree *SumMerkleTree

	txs, err := PrepareLeaves(q)
	if err != nil {
		assert.Error(t, err)
	}

	sumTree = NewSumMerkleTree(txs)

	root := sumTree.GetRoot()

	result.RootHash = fmt.Sprintf("%x", root.Hash)
	result.RootLength = fmt.Sprintf("%x", root.Length)

	leaf := sumTree.Leafs[1]

	result.Begin = fmt.Sprintf("%d", leaf.Begin)
	result.End = fmt.Sprintf("%d", leaf.End)

	proof := sumTree.GetProof(1)
	result.Item = fmt.Sprintf("%x", proof.Item)

	for _, item := range proof.Data {
		result.ItemsLenAndHash = append(result.ItemsLenAndHash, fmt.Sprintf("%x%x", item.Length, item.Hash))
	}

	result.RlpEncoded = fmt.Sprintf("%x", sumTree.GetRlpEncodedProof(1))

	resultJSON, _ := json.Marshal(result)
	err = ioutil.WriteFile("../../contracts/test/result.json", resultJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}

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
			{b1, e1},
			{b2, e2},
		}
		return fmt.Sprint(FillGaps(slices))
	}

	// Slices at the beginning and at the end
	assert.Equal(t, "[{0 10} {10 16777000} {16777000 16777215}]", twoSlices(0, 10, 16777000, 16777215))

	// Two Slices in the middle
	assert.Equal(t, "[{0 100} {100 200} {200 500} {500 600} {600 16777215}]", twoSlices(100, 200, 500, 600))

	// Fill gaps between three slices
	threeSlices := func(b1 uint32, e1 uint32, b2 uint32, e2 uint32, b3 uint32, e3 uint32) string {
		slices := []Slice{
			{b1, e1},
			{b2, e2},
			{b3, e3},
		}
		return fmt.Sprint(FillGaps(slices))
	}

	assert.Equal(t, "[{0 10} {10 200} {200 500} {500 3000} {3000 16777215}]", threeSlices(0, 10, 200, 500, 3000, 16777215))

	// Return just one slice is source collection is empty
	assert.Equal(t, "[{0 16777215}]", fmt.Sprint(FillGaps([]Slice{})))
}
