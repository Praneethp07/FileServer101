package main

import (
	"file-server/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 8080

	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/download", handlers.Download)
	http.HandleFunc("/addUser", handlers.AddNewUser)
	//start the server

	fmt.Printf("server starting at http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
