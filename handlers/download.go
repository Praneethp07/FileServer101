package handlers

import (
	"encoding/json"
	"file-server/models"
	"file-server/utils"
	"fmt"
	"net/http"
	"path/filepath"
)

func Download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "filename parameter is required", http.StatusBadRequest)
		return
	}

	// Read credentials 'creds' field
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
	folderName := utils.GenerateUserFolder(creds.Username, creds.Password, creds.EmailID)
	filePath := filepath.Join(fmt.Sprintf("./STORAGE/%s", folderName), filepath.Clean(filename)+".enc")

	// Set headers BEFORE writing body
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Stream decrypted + decompressed file directly in response
	err := utils.ProcessAndServeFile(w, filePath)
	if err != nil {
		// Cannot safely write error response now: headers and partial body may be sent,
		// so just log or handle accordingly
		http.Error(w, "Failed to process file: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
