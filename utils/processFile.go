package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	// import credentials from models package
)

func ProcessAndStoreFile(file io.Reader, dstPath string) error {
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

func ProcessAndServeFile(w http.ResponseWriter, filePath string) error {
	// Create temp file for decrypted data
	decryptedTempFile, err := os.CreateTemp("", "decrypted_*")
	if err != nil {
		return fmt.Errorf("failed to create temp file for decrypted data: %w", err)
	}
	defer os.Remove(decryptedTempFile.Name())
	defer decryptedTempFile.Close()

	// 1. Decrypt the file into the temp decrypted file
	err = DecryptFile(filePath, decryptedTempFile)
	if err != nil {
		return fmt.Errorf("failed to decrypt file: %w", err)
	}

	// Sync and close decrypted temp file to flush all writes
	if err = decryptedTempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync decrypted temp file: %w", err)
	}
	if err = decryptedTempFile.Close(); err != nil {
		return fmt.Errorf("failed to close decrypted temp file: %w", err)
	}

	// Create temp file for decompressed data
	decompressedTempFile, err := os.CreateTemp("", "decompressed_*")
	if err != nil {
		return fmt.Errorf("failed to create temp file for decompressed data: %w", err)
	}
	defer os.Remove(decompressedTempFile.Name())
	defer decompressedTempFile.Close()

	// 2. Decompress the decrypted temp file into decompressed temp file
	err = DecompressFile(decryptedTempFile.Name(), decompressedTempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to decompress file: %w", err)
	}

	// Rewind decompressed file for reading
	if _, err = decompressedTempFile.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek decompressed temp file: %w", err)
	}

	// 3. Serve the decompressed file content
	_, err = io.Copy(w, decompressedTempFile)
	if err != nil {
		return fmt.Errorf("failed to write decompressed data to response: %w", err)
	}
	
	return nil
}
