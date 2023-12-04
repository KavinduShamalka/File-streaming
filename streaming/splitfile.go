package streaming

import (
	"fmt"
	"io"
	"os"
)

// SplitFile splits a file into chunks of a specified size
func SplitFile(inputFile string, chunkSize int64) ([]string, error) {

	// Open the input file
	file, err := os.Open(inputFile)

	if err != nil {
		return nil, err
	}

	// Close the file, when the function exits
	defer file.Close()

	// Get file info
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
