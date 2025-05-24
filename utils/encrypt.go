package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func EncryptFile(srcPath, destPath string) error {
	key := os.Getenv("ENCRYPTION_KEY")

	if len(key) != 32 {
		return fmt.Errorf("encryption key must be 32 bytes but found %d bytes", len(key))
	}
	fmt.Printf("source file path:%s\n", srcPath)
	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("trying to open source file")
		return err
	}
	defer src.Close()

	//create destination path
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	//create aes cipher
	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	initVector := make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(rand.Reader, initVector); err != nil {
		return err
	}

	//write initialization vector to destination

	if _, err := dst.Write(initVector); err != nil {
		return err
	}
	stream := cipher.NewCTR(cipherBlock, initVector)
	_, err = io.Copy(cipher.StreamWriter{S: stream, W: dst}, src)
	if err == nil {
		log.Printf("Encryption successfull...")
	}
	return err
}
