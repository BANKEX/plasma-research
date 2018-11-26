package utils

import (
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
)

type Item []byte
type Layer []Item

type MerkleTree struct {
	Layers []Layer
}

func NewMerkleTree(data []Item) *MerkleTree {
	tree := new(MerkleTree)
	lambda := func(x Item) Item { return hashFunction(x) }
	tree.Layers = getLayers(Map(lambda, data))
	return tree
}

func hashFunction(data Item) Item {
	return crypto.Keccak256(data)
}

type Tree []Layer

type mapDelegate func(Item) Item

func Map(f mapDelegate, list []Item) []Item {
	ys := make([]Item, len(list))
	for i, x := range list {
		ys[i] = f(x)
	}
	return ys
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

func getLayers(elements []Item) Tree {
	emptyLeveled := hashFunction([]byte{})
	if (len(elements) % 2) > 0 {
		elements = append(elements, emptyLeveled)
	}

	myTree := Tree{elements}
	const maxLevel = 31

	for level := 1; level <= maxLevel; level++ {
		var current []Item
		for i := 0; i < len(myTree[level-1])/2; i++ {
			a := myTree[level-1][i*2]
			b := myTree[level-1][i*2+1]
			hash := hashFunction(concat(a, b))

			current = append(current, hash)
		}

		if (len(current)%2 > 0) && level < maxLevel {
			current = append(current, emptyLeveled)
		}

		emptyLeveled = hashFunction(concat(emptyLeveled, emptyLeveled))

		myTree = append(myTree, current)
	}

	return myTree
}

func (tree *MerkleTree) getRoot() Item {
	return tree.Layers[len(tree.Layers)-1][0]
}

func (tree *MerkleTree) GetHexRoot() Item {
	return tree.getRoot()
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

//getHexProof (idx) {
//  const result = this.getProof(idx);
//  return this.bufArrToHexArr(result);
//}
//

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

//bufIndexOf (el, arr) {
//let hash;
//
//// Convert element to 32 byte hash if it is not one already
//if (el.length != = 32 || !Buffer.isBuffer(el)) {
//hash = this.hashFunction(el);
//} else {
//hash = el;
//}
//
//for (let i = 0; i < arr.length; i++) {
//if (hash.equals(arr[i])) {
//return i;
//}
//}
//
//return -1;
//}
//
//bufArrToHexArr (arr) {
//if (arr.some(el = > !Buffer.isBuffer(el))) {
//throw new Error('Array is not an array of buffers');
//}
//
//return arr.map(el = > '0x' + el.toString('hex'));
//}
//}
//
//function keccak160 (element) {
//return keccak256(element).slice(12,32);
//}
//
//module.exports = {
//MerkleTree,
//keccak160
//};
