package utils

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Item []byte
type Layer []Item
type Tree []Layer

// Declare type since we use Keccak160 in production and Keccak256 in JS tests
type HashFunction func(data Item) Item

type MerkleTree struct {
	Layers []Layer
}

func Keccak256(data Item) Item {
	return crypto.Keccak256(data)
}

func Keccak160(data Item) Item {
	return crypto.Keccak256(data)[12:32]
}

func NewMerkleTree(leafsData []Item, height int, hash HashFunction) *MerkleTree {

	tree := new(MerkleTree)
	var leafs []Item
	for _, data := range leafsData {
		leafs = append(leafs, hash(data))
	}
	tree.Layers = getLayers(leafs, height, hash)

	return tree
}

func getLayers(elements []Item, maxLevel int, hash HashFunction) Tree {
	emptyLeveled := hash([]byte{})
	if (len(elements) % 2) > 0 {
		elements = append(elements, emptyLeveled)
	}

	myTree := Tree{elements}

	for level := 1; level <= maxLevel; level++ {
		var current []Item
		for i := 0; i < len(myTree[level-1])/2; i++ {
			a := myTree[level-1][i*2]
			b := myTree[level-1][i*2+1]
			hash := hash(concat(a, b))

			current = append(current, hash)
		}

		if (len(current)%2 > 0) && level < maxLevel {
			current = append(current, emptyLeveled)
		}

		emptyLeveled = hash(concat(emptyLeveled, emptyLeveled))
		myTree = append(myTree, current)
	}

	return myTree
}

func (tree *MerkleTree) GetRoot() Item {
	return tree.Layers[len(tree.Layers)-1][0]
}

func (tree *MerkleTree) GetHexRoot() Item {
	return tree.GetRoot()
}

func (tree *MerkleTree) GetProof(idx int) []Item {
	lambda := func(proof []Item, layer []Item) []Item {
		pairElement := getPairElement(idx, layer)
		if pairElement != nil {
			proof = append(proof, pairElement)
		}
		idx = idx / 2
		return proof
	}

	return Reduce(tree.Layers, lambda, []Item{})
}

type reduceDelegate func(accumulator []Item, currentValue []Item) []Item
func Reduce(layers []Layer, f reduceDelegate, zero []Item) []Item {

	layersLen := len(layers)
	switch layersLen {
	case 0:
		return zero
	case 1:
		return layers[0]
	}

	out := f(zero, layers[0])
	for i := 1; i < layersLen; i++ {
		out = f(out, layers[i])
	}
	return out
}

func concat(values ...Item) Item {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.Write(s)
	}
	return buffer.Bytes()
}

func (tree *MerkleTree) getHexProof(idx int) string {
	return fmt.Sprintf("%x", tree.GetProof(idx))
}

func getPairElement(idx int, layer []Item) Item {
	pairIdx := idx
	if (idx % 2) == 0 {
		pairIdx = idx + 1
	} else {
		pairIdx = idx - 1
	}

	if pairIdx < len(layer) {
		return layer[pairIdx]
	} else {
		return nil
	}
}
