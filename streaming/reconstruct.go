package streaming

import (
	"io"
	"os"
)

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
