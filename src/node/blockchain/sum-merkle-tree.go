package blockchain

import (
	"fmt"
	. "github.com/BANKEX/plasma-research/src/node/alias"
	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	. "github.com/BANKEX/plasma-research/src/node/utils"
	"sort"
)

type SumTreeRoot struct {
	// We use 24 bit
	Length uint32
	Hash   Uint160
}

type SumTreeNode struct {
	Begin uint32
	End   uint32

	// We use 24 bit of length field
	Length uint32
	Hash   Uint160

	Left   *SumTreeNode
	Right  *SumTreeNode
	Parent *SumTreeNode
}

// Index: Bit index of the element in the tree
// Slice: Slice that stored inside a leaf with corresponding index
// Item: Hash ot the transaction associated with a slice
// Data: List of proof steps
type SumMerkleTreeProof struct {
	Index uint32
	Slice Slice
	Item  Uint160
	Data  []ProofStep
}

type ProofStep struct {
	Length []byte  // 4 bytes
	Hash   Uint160 // 20 bytes
}


func HasIntersection(slices []Slice) error {
	for i := 0; i < len(slices)-1; i++ {
		if slices[i].End > slices[i+1].Begin {
			return fmt.Errorf("slices (%d, %d) and (%d, %d) intersect",
				slices[i].Begin, slices[i].End, slices[i+1].Begin, slices[i+1].End)
		}
	}
	return nil
}

// Use this first when assemble blocks
func PrepareLeaves(transactions []Transaction) ([]*SumTreeNode, error) {

	zeroHash := Keccak160([]byte{})
	slice2transactions := map[Slice]*Transaction{}

	var slices []Slice
	for _, t := range transactions {
		for _, input := range t.Inputs {
			slices = append(slices, input.Slice)
			slice2transactions[input.Slice] = &t
		}
	}

	sort.Slice(slices, func(i, j int) bool {
		return slices[i].Begin < slices[j].Begin
	})


	err := HasIntersection(slices)
	if err != nil{
		return nil, err
	}

	slices = FillGaps(slices)

	var leaves []*SumTreeNode
	for _, slice := range slices {

		// Slices that filling the gaps haven't got a reference to transaction
		tx, hasTx := slice2transactions[slice]
		var txHash = zeroHash
		if hasTx {
			txHash = tx.GetHash()
		}

		leaf := SumTreeNode{
			Begin:  slice.Begin,
			End:    slice.End,
			Hash:   txHash,
			Length: slice.End - slice.Begin,
		}

		leaves = append(leaves, &leaf)
	}
	return leaves, nil
}

type SumMerkleTree struct {
	Root  *SumTreeNode
	Leafs []*SumTreeNode
}

func uint32BE(n uint32) []byte {
	return []byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
}

func concatAndHash(left *SumTreeNode, right *SumTreeNode, hashFunc HashFunction) Uint160 {
	l1, l2 := left.Length, right.Length
	h1, h2 := left.Hash, right.Hash

	d1 := append(uint32BE(l1), uint32BE(l2)...)
	d2 := append(h1, h2...)

	result := append(d1, d2...)
	return hashFunc(result)
}

func NewSumMerkleTree(leafs []*SumTreeNode, hashFunc HashFunction) *SumMerkleTree {
	var tree SumMerkleTree
	tree.Leafs = leafs

	var buckets = tree.Leafs

	// At the end we assign new layer to buckets, so stop when ever we can't merge anymore
	for len(buckets) != 1 {
		// next layer
		var newBuckets []*SumTreeNode

		for len(buckets) > 0 {
			if len(buckets) >= 2 {

				// deque pair from the head
				left, right := buckets[0], buckets[1]
				buckets = buckets[2:]

				length := left.Length + right.Length
				hash := concatAndHash(left, right, hashFunc)

				node := SumTreeNode{
					Hash:   hash,
					Length: length,
				}

				left.Parent = &node
				right.Parent = &node

				left.Right = right
				right.Left = left

				newBuckets = append(newBuckets, &node)

			} else {
				// Pop the last one - goes to next layer as it is
				newBuckets = append(newBuckets, buckets[0])
				buckets = []*SumTreeNode{}
			}
		}
		buckets = newBuckets
	}

	tree.Root = buckets[0]
	return &tree
}

func (tree *SumMerkleTree) GetProof(leafIndex uint32) SumMerkleTreeProof {

	var curr = tree.Leafs[leafIndex]
	var result SumMerkleTreeProof
	result.Slice = Slice{curr.Begin, curr.End}
	result.Item = curr.Hash

	index := uint32(0)
	var proofSteps []ProofStep

	for i := uint(0); curr.Parent != nil; i++ {

		var node *SumTreeNode
		if curr.Right != nil {
			// We are on the left
			node = curr.Right
		} else {
			// We have left node - it means we are at the right
			node = curr.Left
			// set bit in index, if we are at the right
			index |= 1 << i
		}

		// 4 + 20 byte
		// step := append(uint32BE(node.Length), node.Hash...)
		// proofSteps = append(proofSteps, step...)

		step := ProofStep{uint32BE(node.Length), node.Hash}
		proofSteps = append(proofSteps, step)
		curr = curr.Parent
	}

	result.Index = index
	result.Data = proofSteps
	return result
}

func (tree *SumMerkleTree) GetRlpEncodedProof(leafIndex uint32) []byte {
	proof := tree.GetProof(leafIndex)

	var data []byte
	for _, proofItem := range proof.Data {
		data = append(data, proofItem.Length...)
		data = append(data, proofItem.Hash...)
	}

	tmp := struct {
		Index uint32
		Slice Slice
		Item  []byte
		Data  []byte
	}{
		proof.Index,
		proof.Slice,
		proof.Item,
		data,
	}

	rlp, _ := EncodeToRLP(tmp)
	return rlp
}

func (tree *SumMerkleTree) GetRoot() SumTreeRoot {
	r := tree.Root
	return SumTreeRoot{
		r.Length,
		r.Hash,
	}
}

// We use 24 bits to define plasma slices space
// 2^24 - 1 = 0x00FFFFFF
const plasmaLength = 16777215

// Fill plasma range space with Slices, src slices should be sorted first
func FillGaps(src []Slice) []Slice {

	// TODO(artall64): Slice Merge, it doesn't merge a slices even if they are neighbors as I remember such improvement can be useful
	var result []Slice

	pos := uint32(0)
	for i := 0; i <= len(src)-1; i++ {

		item := src[i]

		if pos < item.Begin {
			emptySlice := Slice{
				Begin: pos,
				End:   item.Begin,
			}
			result = append(result, emptySlice)
		}

		result = append(result, item)
		pos = item.End
	}

	if pos != plasmaLength {
		emptySlice := Slice{
			Begin: pos,
			End:   plasmaLength,
		}
		result = append(result, emptySlice)
	}

	return result
}
