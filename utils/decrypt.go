package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

func DecryptFile(srcPath string, w io.Writer) error {
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		return fmt.Errorf("ENCRYPTION_KEY must be 32 bytes")
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	// Read IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(src, iv); err != nil {
		return err
	}
	// Create AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Decrypt
	stream := cipher.NewCTR(block, iv)
	_, err = io.Copy(w, &cipher.StreamReader{S: stream, R: src})
	return err
}
