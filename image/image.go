package image

import (
	"encoding/base64"
	"io"
	"os"
)

func Generate(content, filePath string) error {
	imageBytes := []byte(content)
	err := os.WriteFile(filePath, imageBytes, 0o755)
	return err
}

func GenerateByBytes(content []byte, filePath string) error {
	imageBytes := content
	err := os.WriteFile(filePath, imageBytes, 0o755)
	return err
}

func ToBase64(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// read file content
	imageData, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	// encode to base64
	base64String := base64.StdEncoding.EncodeToString(imageData)
	return base64String, nil
}
