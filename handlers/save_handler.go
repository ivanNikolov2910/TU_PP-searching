package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// SaveHandler handles the saving of search details to a file.
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var searchDetails map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&searchDetails)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("resources/results/search_results_%d.json", time.Now().UnixNano())
	err = saveToFile(filePath, searchDetails)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Search details saved successfully!"})
}

func saveToFile(filePath string, data interface{}) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
