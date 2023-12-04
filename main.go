package main

import (
	"chunk/streaming"
	"fmt"
)

func main() {
	inputFile := "text.png"
	chunkSize := int64(102400) // Set your desired chunk size in bytes

	// Split the file into chunks
	chunkNames, err := streaming.SplitFile(inputFile, chunkSize)
	if err != nil {
		fmt.Println("Error splitting file:", err)
		return
	}

	// Retrieve and concatenate the chunks to reconstruct the original file
	outputFileName := "new.jpg"
	err = streaming.RetrieveChunks(chunkNames, outputFileName)
	if err != nil {
		fmt.Println("Error retrieving chunks:", err)
		return
	}

	fmt.Println("File split and retrieved successfully!")
}
