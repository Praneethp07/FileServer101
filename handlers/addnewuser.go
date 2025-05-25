package handlers

import (
	"encoding/json"
	"file-server/models"
	"file-server/utils"
	"fmt"
	"net/http"
	"os"
)

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	var creds models.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	folderName := utils.GenerateUserFolder(creds.Username, creds.Password, creds.EmailID)
	fmt.Println("foldername:", folderName)
	folder := fmt.Sprintf("STORAGE/%s", folderName)
	err = os.Mkdir(folder, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("Folder %s created...", folder)
	}
}
