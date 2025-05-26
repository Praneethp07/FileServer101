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

	// Parse the multipart form data (limit 10 MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Read the credentials JSON string from the "creds" form field
	credsJSON := r.FormValue("creds")
	if credsJSON == "" {
		http.Error(w, "`creds` field is required", http.StatusBadRequest)
		return
	}

	var creds models.UserCredentials
	if err := json.Unmarshal([]byte(credsJSON), &creds); err != nil {
		http.Error(w, "Failed to parse creds JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the uploaded file from form field "file"
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate folder name for the user based on creds
	folderName := utils.GenerateUserFolder(creds.Username, creds.Password, creds.EmailID)

	// Define base storage path
	baseStoragePath := "./STORAGE"

	// Full path for user folder
	userFolderPath := filepath.Join(baseStoragePath, folderName)

	// Create folder if it doesn't exist
	if err := os.MkdirAll(userFolderPath, 0700); err != nil {
		http.Error(w, "Failed to create user folder: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Destination path for encrypted file (.enc extension appended)
	dstPath := filepath.Join(userFolderPath, handler.Filename+".enc")

	// Process and store the file (compress, encrypt, save)
	if err := utils.ProcessAndStoreFile(file, dstPath); err != nil {
		http.Error(w, "Failed to process and store the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload successfull"))
}
