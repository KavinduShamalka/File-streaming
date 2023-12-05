package main

import (
	"chunk/streaming"
	"crypto/sha256"
	"fmt"
	"os"
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

	inputFile, err := os.Open("test.jpg")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	chunkSize := int64(1024) // Set your desired chunk size in bytes

	// Split the file into chunks
	name, err := streaming.SplitFile(inputFile, chunkSize)
	if err != nil {
		fmt.Println("Error splitting file:", err)
		return
	}

	// fmt.Printf("Size: %v\n", name)

	root := buildMerkleTree(name)
	printTree(root, "")
	fmt.Println("Root Hash:", root.Hash)

	// // Print hash values for each chunk
	// fmt.Println("Hash values for each chunk:")
	// for i, hashValue := range hash {
	// 	fmt.Printf("Chunk %d: %s\n", i+1, hashValue)
	// }

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
