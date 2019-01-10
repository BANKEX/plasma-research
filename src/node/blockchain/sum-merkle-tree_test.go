package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/stretchr/testify/assert"
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

	jsonFile, err := os.Open("sum-tree-sample.json")
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

	resultJSON, _ := json.MarshalIndent(result, "", "  ")

	err = ioutil.WriteFile("../../contracts/test/sample-merkle-proof.json", resultJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}

}

func TestPrepareLeaves(t *testing.T) {

	// First check

	var fTxList []Transaction
	var fTx Transaction

	// 2 Inputs with intersected slices

	fTxInput1 := Input{Output: Output{
		Slice: Slice{10, 20},
	}}

	fTxInput2 := Input{Output: Output{
		Slice: Slice{14, 28},
	}}

	//

	fTx.Inputs = append(fTx.Inputs, fTxInput1)
	fTx.Inputs = append(fTx.Inputs, fTxInput2)

	fTxList = append(fTxList, fTx)

	_, fErr := PrepareLeaves(fTxList)
	fExpected := fmt.Errorf("slices (%d, %d) and (%d, %d) intersect",
		fTxInput1.Slice.Begin, fTxInput1.Slice.End, fTxInput2.Slice.Begin, fTxInput2.Slice.End)

	assert.Equal(t, fExpected, fErr)

	// Second check
	var sTxList []Transaction
	var sTx Transaction

	// 3 Inputs with 2 intersected slices (sTxInput2 - sTxInput3)

	sTxInput1 := Input{Output: Output{
		Slice: Slice{4, 7},
	}}

	sTxInput2 := Input{Output: Output{
		Slice: Slice{8, 19},
	}}

	sTxInput3 := Input{Output: Output{
		Slice: Slice{12, 31},
	}}

	//

	sTx.Inputs = append(sTx.Inputs, sTxInput1)
	sTx.Inputs = append(sTx.Inputs, sTxInput2)
	sTx.Inputs = append(sTx.Inputs, sTxInput3)

	sTxList = append(sTxList, sTx)

	_, sErr := PrepareLeaves(sTxList)

	sExpected := fmt.Errorf("slices (%d, %d) and (%d, %d) intersect",
		sTxInput2.Slice.Begin, sTxInput2.Slice.End, sTxInput3.Slice.Begin, sTxInput3.Slice.End)

	assert.Equal(t, sExpected, sErr)

	// Third check

	var tTxList []Transaction
	var tTx Transaction

	// 3 Inputs with 3 intersected slices

	tTxInput1 := Input{Output: Output{
		Slice: Slice{2, 7},
	}}

	tTxInput2 := Input{Output: Output{
		Slice: Slice{4, 10},
	}}

	tTxInput3 := Input{Output: Output{
		Slice: Slice{5, 30},
	}}

	tTx.Inputs = append(tTx.Inputs, tTxInput1)
	tTx.Inputs = append(tTx.Inputs, tTxInput2)
	tTx.Inputs = append(tTx.Inputs, tTxInput3)

	tTxList = append(tTxList, tTx)

	_, tErr := PrepareLeaves(tTxList)
	tExpected := fmt.Errorf("slices (%d, %d) and (%d, %d) intersect",
		tTxInput1.Slice.Begin, tTxInput1.Slice.End, tTxInput2.Slice.Begin, tTxInput2.Slice.End)

	assert.Equal(t, tExpected, tErr)

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
