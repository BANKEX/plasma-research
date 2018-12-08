package blockchain

import (
	a "../alias"
	"sort"
)
import u "../utils"
import s "../../plasmautils/slice"

type SumTreeNode struct {
	// We use 24 bit
	Length uint32
	Hash   a.Uint160

	isLeft  bool
	isRight bool

	Left   *SumTreeNode
	Right  *SumTreeNode
	Parent *SumTreeNode
}

// Use this first when assemble blocks
func PrepareLeafs(transactions []Transaction) []SumTreeNode {

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

	var leafs []SumTreeNode
	for idx, slice := range slices {
		leaf := SumTreeNode{
			Hash:    slice2transactions[slice].GetHash(),
			Length:  slice.End - slice.Begin,
			isLeft:  idx%2 == 0,
			isRight: idx%1 == 0,
		}

		leafs = append(leafs, leaf)
	}

	// TODO: look do we need padding at that moment
	if len(leafs)%2 == 1 {
		emptyLeaf := SumTreeNode{
			Hash:    u.Keccak160([]byte{}), // Hash from empty byte array
			Length:  0,
			isLeft:  false,
			isRight: true,
		}
		leafs = append(leafs, emptyLeaf)
	}

	for i := 0; i < len(leafs); i += 2 {
		left := leafs[i]
		right := leafs[i+1]

		left.isLeft = true
		left.isRight = false
		left.Left = nil
		left.Right = &right

		right.isLeft = false
		right.isRight = true
		right.Left = &left
		right.Right = nil
	}

	return leafs
}

type SumMerkleTree struct {
	Root  SumTreeNode
	Leafs []SumTreeNode
}

func NewSumMerkleTree(leafs []SumTreeNode) *SumMerkleTree {
	var tree SumMerkleTree
	tree.Leafs = leafs

	// Root

	return &tree
}

func (tree *SumMerkleTree) GetProof(leafIndex uint32) []byte {
	return []byte{0x0}
	// return tree.Root.Length, tree.Root.Hash
}

//func (tree *SumMerkleTree) GetRoot() (length uint32, hash a.Uint160) {
//	return tree.Root.Length, tree.Root.Hash
//}

func (tree *SumMerkleTree) GetRoot() SumTreeNode {
	return tree.Root
}
