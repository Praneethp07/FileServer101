package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"file-server/models" // import credentials from models package
)

func ProcessAndStoreFile(file io.Reader, dstPath string, creds models.UserCredentials) error {
	// Step 0: Create temp dir for intermediate files
	tempDir, err := os.MkdirTemp("", "procfile")
	if err != nil {
		return fmt.Errorf("unable to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir) // clean up temp files

	// Step 1: Save uploaded data to temp source file
	tempSrcFile := filepath.Join(tempDir, "uploaded.tmp")
	srcFile, err := os.Create(tempSrcFile)
	if err != nil {
		return fmt.Errorf("unable to create temp source file: %w", err)
	}
	_, err = io.Copy(srcFile, file)
	srcFile.Close()
	if err != nil {
		return fmt.Errorf("error saving uploaded file: %w", err)
	}

	// Step 2: Compress temp source file to another temp file
	tempCompressedFile := filepath.Join(tempDir, "compressed.gz")
	err = CompressFile(tempSrcFile, tempCompressedFile)
	if err != nil {
		return fmt.Errorf("compression failed: %w", err)
	}

	// Step 3: Encrypt compressed file to destination path
	err = EncryptFile(tempCompressedFile, dstPath)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}

	// Step 4: Done successfully
	return nil
}
