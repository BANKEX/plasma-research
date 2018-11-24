package main

import "fmt"
import u "../utils"

func main() {
	a := []byte{1}
	b := []byte{2}
	c := []byte{3}
	d := []byte{7}

	tree := u.NewMerkleTree([]u.Item{a, b, c, d})
	fmt.Printf("%x", tree.GetHexRoot())
}
