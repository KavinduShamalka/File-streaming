package merkle

import (
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	Hash  string
	Left  *MerkleNode
	Right *MerkleNode
}

func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func buildMerkleTree(data []string) *MerkleNode {

	if len(data) == 1 {
		return &MerkleNode{Hash: calculateHash(data[0]), Left: nil, Right: nil}
	}

	mid := len(data) / 2
	left := buildMerkleTree(data[:mid])
	right := buildMerkleTree(data[mid:])
	return &MerkleNode{Hash: calculateHash(left.Hash + right.Hash), Left: left, Right: right}
}

func printTree(node *MerkleNode, indent string) {
	if node != nil {
		fmt.Println(indent+"Hash:", node.Hash)
		if node.Left != nil {
			printTree(node.Left, indent+"  ")
		}
		if node.Right != nil {
			printTree(node.Right, indent+"  ")
		}
	}
}

func main() {

	data := []string{"a", "b", "c"}

	fmt.Println(data)
	mid := len(data) / 2

	fmt.Println("Mid: ", mid)

	fmt.Println(data[:mid])
	fmt.Println(data[mid:])

	root := buildMerkleTree(data)
	printTree(root, "")
	fmt.Println("Root Hash:", root.Hash)
}
