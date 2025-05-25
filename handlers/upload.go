package handlers

import (
	"encoding/json"
	"file-server/models"
	"file-server/utils"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON credentials from request body
	var creds models.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Body is not formatted properly", http.StatusBadRequest)
		return
	}

	// Parse the multipart form data (limit 10 MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form, form field "file"
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate user folder name based on credentials
	folderName := utils.GenerateUserFolder(creds.Username, creds.Password, creds.EmailID)

	// Define base path for storage
	baseStoragePath := "./STORAGE"

	// Build full user folder path
	userFolderPath := filepath.Join(baseStoragePath, folderName)

	// Create folder if it doesn't exist
	if err := os.MkdirAll(userFolderPath, 0700); err != nil {
		http.Error(w, "Failed to create user folder: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Destination path for encrypted file, append ".enc" extension
	dstPath := filepath.Join(userFolderPath, handler.Filename+".enc")

	// Process (compress, encrypt, save) the uploaded file
	err = utils.ProcessAndStoreFile(file, dstPath, creds)
	if err != nil {
		http.Error(w, "Failed to process and store file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded and processed successfully"))
}
