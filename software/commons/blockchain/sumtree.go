package blockchain

import a "../alias"
import u "../utils"

type SumTreeNode struct {
	// We use 24 bit
	Length uint32
	Hash   a.Uint160
}

// TODO should return linear sequence of segments that covers Plasma address space
// If segment affect by transaction from given list it will assign it hash to leaf otherwise it will assign hash(0)
func PrepareLeafs(transactions []Transaction) []SumTreeNode {

	tmp := SumTreeNode{
		PlasmaRangeSpace,
		u.Keccak160([]byte{0x0, 0x0}),
	}

	return []SumTreeNode{tmp}
}

type SumMerkleTree struct {
}

func NewSumMerkleTree(leafs []SumTreeNode) *SumMerkleTree {
	var tree SumMerkleTree
	return &tree
}

func (tree *SumMerkleTree) GetRoot() SumMerkleNode {
	return SumMerkleNode{100, a.Uint160{0x0}}
}
