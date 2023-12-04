package main

import (
	"fmt"
	"io"
	"os"
)

// SplitFile splits a file into chunks of a specified size
func SplitFile(inputFile string, chunkSize int64) ([]string, error) {
	file, err := os.Open(inputFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	chunkNames := make([]string, 0)

	for i := int64(0); i < fileSize; i += chunkSize {
		chunkName := fmt.Sprintf("%s_chunk%d", inputFile, i/chunkSize+1)
		chunkFile, err := os.Create(chunkName)
		if err != nil {
			return nil, err
		}

		// Copy the chunkSize bytes from the original file to the chunk file
		_, err = io.CopyN(chunkFile, file, chunkSize)
		if err != nil && err != io.EOF {
			return nil, err
		}

		chunkFile.Close()
		chunkNames = append(chunkNames, chunkName)
	}

	return chunkNames, nil
}

// RetrieveChunks retrieves and concatenates chunks to reconstruct the original file
func RetrieveChunks(chunkNames []string, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, chunkName := range chunkNames {
		chunkFile, err := os.Open(chunkName)
		if err != nil {
			return err
		}
		defer chunkFile.Close()

		// Copy the content of each chunk to the output file
		_, err = io.Copy(outputFile, chunkFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	inputFile := "text.png"
	chunkSize := int64(102400) // Set your desired chunk size in bytes

	// Split the file into chunks
	chunkNames, err := SplitFile(inputFile, chunkSize)
	if err != nil {
		fmt.Println("Error splitting file:", err)
		return
	}

	// Retrieve and concatenate the chunks to reconstruct the original file
	outputFileName := "new.jpg"
	err = RetrieveChunks(chunkNames, outputFileName)
	if err != nil {
		fmt.Println("Error retrieving chunks:", err)
		return
	}

	fmt.Println("File split and retrieved successfully!")
}
