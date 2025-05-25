package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not formatted properly", http.StatusBadRequest)
		return
	}
	response := []byte("File is being downloaded....")
	fmt.Printf("Request body:%s\n", string(bodyBytes))

	w.Write(response)
}
