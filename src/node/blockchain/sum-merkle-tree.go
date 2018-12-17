package blockchain

import (
	"fmt"
	"sort"

	. "github.com/BANKEX/plasma-research/src/node/alias"
	. "github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type SumMerkleNode struct {
	// We use 24 bit
	Length uint32
	Hash   Uint160
}

type SumMerkleTree struct {
	NodeList []SumMerkleNode
}

type SumMerkleProof struct {
	Index    uint32
	Slice    Slice
	Hash     Uint160
	NodeList []SumMerkleNode
}

// func uint160ToUint256(Uint160 n) Uint256 {
// 	res := make(byte, 256)
// 	res[96:256] = n[:]
// 	return res
// }

func (t SumMerkleTree) GetRoot() SumMerkleNode {
	return t.NodeList[0]
}

func (t SumMerkleTree) GetLeaves() []SumMerkleNode {
	return t.NodeList[1<<16-1 : 1<<17-1]
}

func (t SumMerkleTree) MerkleProof(Index uint32) SumMerkleProof {
	res := *new(SumMerkleProof)
	res.Index = Index
	left := uint32(0)

	curCell := 1<<16 - 1 + Index
	length := t.NodeList[curCell].Length
	res.Hash = t.NodeList[curCell].Hash

	for curCell > 0 {
		if (curCell & 1) == 0 {
			res.NodeList = append(res.NodeList, t.NodeList[curCell-1])
			left += t.NodeList[curCell-1].Length
		} else {
			res.NodeList = append(res.NodeList, t.NodeList[curCell+1])
		}
		curCell = (curCell - 1) >> 1
	}
	res.Slice = Slice{left, left + length}
	return res
}

func uint32BE(n uint32) []byte {
	return []byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
}

type OwnedSlice struct {
	Slice
	TxHash Uint160
}

func Hash(a, b Uint160) Uint160 {
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write(append(a, b...))
	buf = hash.Sum(buf)
	return buf
}

func Hash4(x1, x2, x3, x4 []byte) Uint160 {
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write(append(append(append(x1, x2...), x3...), x4...))
	buf = hash.Sum(buf)
	return buf
}

// TODO: check, that slices are conflicting together.
// TODO: join slices for same transactions
func NewSumMerkleTree(txs []Transaction) (*SumMerkleTree, error) {
	nleaves := 1 << 16
	nnodes := nleaves*2 - 1
	t := new(SumMerkleTree)
	nullhash := make([]byte, 160)
	t.NodeList = make([]SumMerkleNode, nnodes)
	leaves := make([]OwnedSlice, 0)

	for _, tx := range txs {
		for _, in := range tx.Inputs {
			leaves = append(leaves, OwnedSlice{in.Slice, tx.GetHash()})
		}
	}
	sort.Slice(leaves, func(i, j int) bool { return leaves[i].Begin < leaves[j].Begin })

	for i := 0; i < len(leaves)-1; i++ {
		if leaves[i].End >= leaves[i+1].Begin {
			return nil, fmt.Errorf("slices (%d, %d) and (%d, %d) intersect", leaves[i].Begin, leaves[i].End, leaves[i+1].Begin, leaves[i+1].End)
		}
	}

	cursorCell := nleaves - 1
	cursorLeft := uint32(0)
	for _, leaf := range leaves {
		if leaf.Begin > cursorLeft {
			t.NodeList[cursorCell] = SumMerkleNode{leaf.Begin - cursorLeft, nullhash[:]}
			cursorCell++
		}
		t.NodeList[cursorCell] = SumMerkleNode{leaf.End - leaf.Begin, leaf.TxHash}
		cursorCell++
		cursorLeft = leaf.End
	}

	t.NodeList[cursorCell] = SumMerkleNode{1<<24 - cursorLeft, nullhash[:]}
	cursorCell++
	for cursorCell < nnodes {
		t.NodeList[cursorCell] = SumMerkleNode{0, nullhash[:]}
		cursorCell++
	}

	for i := nleaves - 2; i >= 0; i-- {
		t.NodeList[i].Length = t.NodeList[2*i+1].Length + t.NodeList[2*i+2].Length
		t.NodeList[i].Hash = Hash4(uint32BE(t.NodeList[2*i+1].Length), uint32BE(t.NodeList[2*i+2].Length), t.NodeList[2*i+1].Hash, t.NodeList[2*i+1].Hash)
	}
	return t, nil
}
