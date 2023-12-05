package main

import (
	"chunk/streaming"
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
	inputFile := "test.jpg"
	chunkSize := int64(102400) // Set your desired chunk size in bytes

	// Split the file into chunks
	chunkNames, err := streaming.SplitFile(inputFile, chunkSize)
	if err != nil {
		fmt.Println("Error splitting file:", err)
		return
	}

	root := buildMerkleTree(chunkNames)
	printTree(root, "")
	fmt.Println("Root Hash:", root.Hash)

	// root := merkle.MerkleNode {

	// }

	// Retrieve and concatenate the chunks to reconstruct the original file
	// outputFileName := "new.jpg"
	// err = streaming.RetrieveChunks(chunkNames, outputFileName)
	// if err != nil {
	// 	fmt.Println("Error retrieving chunks:", err)
	// 	return
	// }

	// fmt.Println("File split and retrieved successfully!")

}
