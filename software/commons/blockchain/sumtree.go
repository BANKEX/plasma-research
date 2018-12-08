package blockchain

import (
	a "../alias"
	"encoding/binary"
	"sort"
)
import u "../utils"
import s "../../plasmautils/slice"

type SumTreeNode struct {
	// We use 24 bit
	Length uint32
	Hash   a.Uint160

	Left   *SumTreeNode
	Right  *SumTreeNode
	Parent *SumTreeNode
}

// Use this first when assemble blocks
func PrepareLeaves(transactions []Transaction) []*SumTreeNode {

	zeroHash := u.Keccak160([]byte{})
	slice2transactions := map[s.Slice]*Transaction{}

	var slices []s.Slice
	for _, t := range transactions {
		for _, input := range t.Inputs {
			slices = append(slices, input.Slice)
			slice2transactions[input.Slice] = &t
		}
	}

	sort.Slice(slices, func(i, j int) bool {
		return slices[i].Begin < slices[j].Begin
	})

	slices = FillGapsWithSlices(slices)

	var leafs []*SumTreeNode
	for _, slice := range slices {

		tx, hasTx := slice2transactions[slice]
		var txHash = zeroHash
		if hasTx {
			txHash = tx.GetHash()
		}

		leaf := SumTreeNode{
			Hash:   txHash,
			Length: slice.End - slice.Begin,
		}

		leafs = append(leafs, &leaf)
	}

	//// Padding
	//if len(leafs)%2 == 1 {
	//	emptyLeaf := SumTreeNode{
	//		Hash:   u.Keccak160([]byte{}), // Hash from empty byte array
	//		Length: 0,
	//	}
	//	leafs = append(leafs, &emptyLeaf)
	//}

	return leafs
}

type SumMerkleTree struct {
	Root  *SumTreeNode
	Leafs []*SumTreeNode
}

// TODO: check that way is compatible with soidity
// Uint to bytes
func u2b(value uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, value)
	return b
}

func concatAndHash(left *SumTreeNode, right *SumTreeNode) a.Uint160 {
	l1, l2 := left.Length, right.Length
	h1, h2 := left.Hash, right.Hash

	d1 := append(u2b(l1), u2b(l2)...)
	d2 := append(h1, h2...)

	result := append(d1, d2...)
	return u.Keccak160(result)
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

func (tree *SumMerkleTree) GetProof(leafIndex uint32) []byte {

	index := uint32(0)
	var curr = tree.Leafs[leafIndex]
	var proofSteps []byte

	for i := uint(0); curr.Parent != nil; i++ {

		var node *SumTreeNode
		if curr.Right != nil {
			node = curr.Right

		} else {
			// We have left node - it means we are at the right
			node = curr.Left
			// set bit in index
			index |= (1 << i)
		}

		// 4 + 20 byte
		step := append(u2b(node.Length), node.Hash...)
		proofSteps = append(proofSteps, step...)

		curr = curr.Parent
	}

	result := append(u2b(index), proofSteps...)
	return result
}

func (tree *SumMerkleTree) GetRoot() *SumTreeNode {
	return tree.Root
}

// 2^24 - 1
const plasmaLength = 16777215

// Fill plasma range space with segments, src slices should be sorted
// 1) It fill the gaps between segments with empty slices
// TODO: NOTE: It doesn't merge a slices even if they are neighbors - as I remember such feature can speedup plasma
func FillGapsWithSlices(src []s.Slice) []s.Slice {

	var result []s.Slice

	for i := 0; i <= len(src); i++ {

		// TODO: Should be out of the cycle
		// Add initial zero slice if needed
		if i == 0 && src[i].Begin != 0 {
			emptySlice := s.Slice{Begin: 0, End: src[i].Begin}
			result = append(result, emptySlice)
		}

		// TODO: Should be out of the cycle
		// Handle last slice and add final empty slice if needed
		if i == len(src)-1 {
			result = append(result, src[i])

			if src[i].End != plasmaLength {
				emptySlice := s.Slice{
					Begin: src[i].End,
					End:   plasmaLength,
				}
				result = append(result, emptySlice)
			}
			return result // !!! Attention, The actual end of the function is Here !!!
		}

		el := src[i]
		nextEl := src[i+1]

		result = append(result, el)
		if nextEl.Begin-el.End > 1 {
			emptySlice := s.Slice{
				Begin: el.End,
				End:   nextEl.Begin,
			}
			result = append(result, emptySlice)
		}
	}

	return result
}
