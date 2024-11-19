package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	sreacher "parallel.searching/v2/parallel_search"
	"parallel.searching/v2/parser"
)

// ExecSearch executes the search based on the provided search type and key.
func ExecSearch(w http.ResponseWriter, ctx *Context, threadCount int) {
	var startTime time.Time
	var result bool

	switch ctx.SearchType {
	case LINEAR:
		startTime = time.Now()
		result = sreacher.ParallelLinearSearch(ctx.Data, ctx.Key, threadCount)
	case BINARY:
		startTime = time.Now()
		result = sreacher.ParallelBinarySearch(ctx.Data, ctx.Key, threadCount)
	case HASH:
		parseDataStartTime := time.Now()
		set := parser.ParseArrayToSet(ctx.Data)
		parseDataTime := time.Since(parseDataStartTime)
		log.Println("Time to parse data to hash set: ", parseDataTime.Seconds(), "s")
		ctx.Size = len(set)
		startTime = time.Now()
		result = sreacher.ParallelHashSearch(set, ctx.Key, threadCount)
	case TREE:
		parseDataStartTime := time.Now()
		root := parser.ParseArrayToNode(ctx.Data)
		parseDataTime := time.Since(parseDataStartTime)
		log.Println("Time to parse data to binary tree: ", parseDataTime.Seconds(), "s")
		startTime = time.Now()
		result = sreacher.ParallelTreeSearch(root, ctx.Key)
		fmt.Println("Done")
	default:
		sendError(w, "Invalid algorithm selection")
		return
	}

	searchTime := time.Since(startTime).Seconds()

	response := SearchResult{
		Result:      result,
		SearchType:  ctx.SearchType,
		SearchedKey: ctx.Key,
		ListSize:    ctx.Size,
		Data:        parser.ParseArrayToString(ctx.Data),
		Range:       fmt.Sprintf("[0 - %v]", ctx.DataRange),
		SearchTime:  fmt.Sprintf("%.0f ns", searchTime*1000000000),
	}

	_ = json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, errorMessage string) {
	response := map[string]string{"error": errorMessage}
	_ = json.NewEncoder(w).Encode(response)
}
