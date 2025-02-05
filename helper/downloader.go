package helper

import (
	"crypto/rand"
	"os"
)

func GenerateLargeFile(filePath string, sizeInMB int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a random large file
	data := make([]byte, 1024*1024) // 1MB
	rand.Read(data)
    for i := 0; i < sizeInMB; i++ {
		_, err := file.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
