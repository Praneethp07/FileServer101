package utils

import (
	"compress/gzip"
	"io"
	"os"
)

func DecompressFile(compressedFilePath, dstFilePath string) error {
	compressedFile, err := os.Open(compressedFilePath)
	if err != nil {
		return err
	}
	defer compressedFile.Close()

	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, gzipReader)
	return err
}
