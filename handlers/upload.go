package handlers

import (
	"file-server/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not formatted properly", http.StatusBadRequest)
		return
	}
	// response := []byte("File is being Uploaded....")
	fmt.Printf("Request body:%s\n", string(bodyBytes))
	srcPath := "/mnt/c/Users/prap/Desktop/PersonalLearngings/FILESTORAGE/FILE-SERVER/Storage/files/file1.txt"
	destPath := "/mnt/c/Users/prap/Desktop/PersonalLearngings/FILESTORAGE/FILE-SERVER/Storage/files/encryptedFiles/encryptedfile.txt"
	err = utils.EncryptFile(srcPath, destPath)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	encryptedFile, err := os.Open(destPath)
	if err != nil {
		fmt.Println(err)
	}
	encryptedContent, err := io.ReadAll(encryptedFile)
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(string(encryptedContent)))
}
