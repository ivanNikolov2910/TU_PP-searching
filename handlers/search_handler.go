package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"parallel.searching/v2/parser"
	"parallel.searching/v2/search"
)

// SearchHandler handles the search request and serves the index page.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, filepath.Join("resources", "index.html"))
		return
	}

	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		searchType := search.SearchType(r.FormValue("algorithm"))
		dataType := search.DataType(r.FormValue("dataType"))

		data, dataRange, size, err := prepareData(w, r, dataType)
		if err != nil {
			return
		}

		threadCountStr := r.FormValue("threadCount")
		threadCount, err := strconv.Atoi(strings.TrimSpace(threadCountStr))
		if err != nil || threadCount < 1 {
			sendError(w, "Invalid thread count value")
			return
		}

		keyStr := r.FormValue("key")
		key, err := strconv.Atoi(strings.TrimSpace(keyStr))
		if err != nil {
			sendError(w, "Invalid key value")
			return
		}
		ctx := &search.Context{
			SearchType: searchType,
			Data:       data,
			DataRange:  dataRange,
			Size:       size,
			Key:        key,
		}
		search.ExecSearch(w, ctx, threadCount)
	}
}

func prepareData(w http.ResponseWriter, r *http.Request, dataType search.DataType) ([]int, int, int, error) {
	if dataType == search.MANUAL {
		data := r.FormValue("manualData")
		if data == "" {
			sendError(w, "Manual data is empty")
			return nil, 0, 0, fmt.Errorf("manual data is empty")
		}
		parsedData := parser.ParseStringToArray(data)
		return parsedData, 0, len(parsedData), nil
	}

	// Handle generated data
	maxElementStr := r.FormValue("maxElement")
	maxElement, err := strconv.Atoi(strings.TrimSpace(maxElementStr))
	if err != nil || maxElement <= 0 {
		sendError(w, "Max element could not be parsed")
		return nil, 0, 0, fmt.Errorf("invalid max element")
	}

	sizeStr := r.FormValue("size")
	size, err := strconv.Atoi(strings.TrimSpace(sizeStr))
	if err != nil {
		sendError(w, "Invalid data input")
		return nil, 0, 0, fmt.Errorf("invalid data input")
	}

	generateDataStartTime := time.Now()
	data, err := search.GetDataByType(dataType, size, maxElement)
	generateDataTime := time.Since(generateDataStartTime)
	log.Println("Time to generate data: ", generateDataTime.Seconds(), "s")
	if err != nil {
		sendError(w, fmt.Sprintf("Invalid data input: %v", err))
		return nil, 0, 0, err
	}
	return data, maxElement, size, nil
}

func sendError(w http.ResponseWriter, errorMessage string) {
	response := map[string]string{"error": errorMessage}
	_ = json.NewEncoder(w).Encode(response)
}
