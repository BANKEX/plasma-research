package blockchain

import (
	"fmt"
	"sort"

	. "github.com/BANKEX/plasma-research/src/node/alias"
	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	. "github.com/BANKEX/plasma-research/src/node/utils"
)

type SumTreeRoot struct {
	// We use 24 bit
	Length uint32
	Hash   Uint160
}

type SumTreeNode struct {
	//
	Begin uint32
	End   uint32

	// We use 24 bit of length field
	Length uint32
	Hash   Uint160

	Left   *SumTreeNode
	Right  *SumTreeNode
	Parent *SumTreeNode
}

type ProofStep struct {
	Length []byte  // 4 bytes
	Hash   Uint160 // 20 bytes
}

// Proof has the same structure as proof in Solidity
type SumMerkleTreeProof struct {
	Index uint32
	Slice Slice
	Item  Uint160
	Data  []ProofStep
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

	for i := 0; i < len(slices)-1; i++ {
		if slices[i].End >= slices[i+1].Begin {
			return nil, fmt.Errorf("slices (%d, %d) and (%d, %d) intersect",
				slices[i].Begin, slices[i].End, slices[i+1].Begin, slices[i+1].End)
		}
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

func concatAndHash(left *SumTreeNode, right *SumTreeNode) Uint160 {
	l1, l2 := left.Length, right.Length
	h1, h2 := left.Hash, right.Hash

	d1 := append(uint32BE(l1), uint32BE(l2)...)
	d2 := append(h1, h2...)

	result := append(d1, d2...)
	return Keccak160(result)
}

func NewSumMerkleTree(leafs []*SumTreeNode) *SumMerkleTree {
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
				hash := concatAndHash(left, right)

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
		//step := append(uint32BE(node.Length), node.Hash...)
		//proofSteps = append(proofSteps, step...)

		step := ProofStep{uint32BE(node.Length), node.Hash}
		proofSteps = append(proofSteps, step)
		curr = curr.Parent
	}

	result.Index = index
	result.Data = proofSteps
	return result
}

//func (tree *SumMerkleTree) GetProof(leafIndex uint32) []byte {
//
//	index := uint32(0)
//	var curr = tree.Leafs[leafIndex]
//	var proofSteps []byte
//
//	for i := uint(0); curr.Parent != nil; i++ {
//
//		var node *SumTreeNode
//		if curr.Right != nil {
//			node = curr.Right
//
//		} else {
//			// We have left node - it means we are at the right
//			node = curr.Left
//			// set bit in index
//			index |= 1 << i
//		}
//
//		// 4 + 20 byte
//		step := append(uint32BE(node.Length), node.Hash...)
//		proofSteps = append(proofSteps, step...)
//
//		curr = curr.Parent
//	}
//
//	result := append(uint32BE(index), proofSteps...)
//	return result
//}

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

	// TODO: Slice Merge
	// It doesn't merge a slices even if they are neighbors
	// as I remember such improvement can speedup plasma

	var result []Slice

	if src[0].Begin != 0 {
		emptySlice := Slice{Begin: 0, End: src[0].Begin}
		result = append(result, emptySlice)
	}

	// if src[0].Begin == 0 && src[0].End != src[1].Begin{
	// 	emptySlice := Slice{Begin: src[0].End, End: src[1].Begin}
	// 	result = append(result, emptySlice)
	// }

	for i := 0; i <= len(src); i++ {

		if i == len(src)-1 {
			result = append(result, src[i])
			break
		}

		el := src[i]
		nextEl := src[i+1]

		result = append(result, el)
		if nextEl.Begin-el.End >= 1 {
			emptySlice := Slice{
				Begin: el.End,
				End:   nextEl.Begin,
			}
			result = append(result, emptySlice)
		}
	}

	if result[len(result)-1].End != plasmaLength {
		emptySlice := Slice{
			Begin: src[len(src)-1].End,
			End:   plasmaLength,
		}
		result = append(result, emptySlice)
	}

	return result
}
