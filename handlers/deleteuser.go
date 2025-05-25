package handlers

import (
	"encoding/json"
	"file-server/models"
	"file-server/utils"
	"fmt"
	"net/http"
	"os"
)

func isPathAvailable(folderPath string) bool {
	info, err := os.Stat(folderPath)
	fmt.Println("stats:", info)
	return os.IsNotExist(err)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var creds models.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	folderName := utils.GenerateUserFolder(creds.Username, creds.Password, creds.EmailID)
	folderPath := fmt.Sprintf("STORAGE/%s", folderName)
	if isPathAvailable(folderPath) {
		//delete user
		if err := os.RemoveAll(folderPath); err != nil {
			fmt.Println(err)
			return
		}

	} else {
		http.Error(w, "this user is not available..", http.StatusBadRequest)
	}

}
