package engine

import (
	"fmt"
	"net/http"
	"parallel.searching/v2/handlers"
)

// WebRun starts the server and sets up the routes.
func WebRun() {
	http.HandleFunc("/", handlers.SearchHandler)
	http.HandleFunc("/save", handlers.SaveHandler)
	fmt.Println("Starting server at http://localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Server could not start properly: %v", err)
	}
}
