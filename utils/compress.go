package utils

import (
	"compress/gzip"
	"io"
	"os"
)

func CompressFile(srcFilePath, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	gzipWriter := gzip.NewWriter(dstFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, srcFile)
	return err
}
